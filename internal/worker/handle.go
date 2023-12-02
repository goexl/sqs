package worker

import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
	"github.com/goexl/log"
	"github.com/goexl/sqs/internal/constant"
	"github.com/goexl/sqs/internal/internal/exception"
	"github.com/goexl/sqs/internal/kernel"
	"github.com/goexl/sqs/internal/message"
	"github.com/goexl/sqs/internal/param"
)

type Handle struct {
	logger  log.Logger
	receive *param.Receive
	param   *param.Handle
	counts  map[string]int
}

func NewHandle(logger log.Logger, receive *param.Receive, param *param.Handle) *Handle {
	return &Handle{
		logger:  logger,
		receive: receive,
		param:   param,
		counts:  make(map[string]int),
	}
}

func (h *Handle) Start(ctx context.Context, handler message.Handler[any]) (err error) {
	cancel, callback := context.WithCancel(ctx)
	h.receive.Cancel(callback)
	if url, ue := h.receive.Url(cancel, h.receive.Base); nil != ue {
		err = ue
	} else {
		err = h.do(cancel, url, handler)
	}

	return

}

func (h *Handle) do(ctx context.Context, url *string, handler message.Handler[any]) (err error) {
	for {
		if out, re := h.receive.Do(ctx, url); nil != re && !h.receive.Exited() {
			h.logger.Warn("收取消息出错", field.New("url", url), field.Error(re))
		} else if nil != out { // 并行消费，加快消费速度
			for _, msg := range out.Messages {
				cloned := msg
				go func() {
					_ = h.handle(ctx, url, &cloned, handler)
				}()
			}
		}

		// 检查是否退出
		if h.receive.Exited() {
			break
		}
	}

	return
}

func (h *Handle) handle(ctx context.Context, url *string, msg *types.Message, handler message.Handler[any]) (err error) {
	if _, ok := msg.MessageAttributes[constant.Runtime]; ok {
		err = h.checkDelay(ctx, url, msg, handler)
	} else {
		err = h.deal(ctx, url, msg, handler)
	}

	return
}

func (h *Handle) deal(ctx context.Context, url *string, msg *types.Message, handler message.Handler[any]) (err error) {
	defer h.cleanup(ctx, url, msg, &err)

	for times := 0; times < h.param.Times; times++ {
		err = h.process(ctx, msg, handler)
		if nil == err {
			break
		} else {
			time.Sleep(h.param.Interval)
		}
	}

	return
}

func (h *Handle) checkDelay(
	ctx context.Context,
	url *string,
	msg *types.Message,
	handler message.Handler[any],
) (err error) {
	now := time.Now()
	sent, runtime := h.getTime(msg)
	if runtime.Before(now) { // 已经过了执行时间，处理消息
		err = h.deal(ctx, url, msg, handler)
	} else if attributes, gae := h.receive.GetAttributes(ctx, url); nil != gae {
		err = gae
	} else if attributes.Invalidate(sent) { // 消息失效，重新发送一个全新消息
		err = h.renew(ctx, url, msg)
	} else { // 改变可见性，等待下一次消费
		visibility := attributes.Visibility()
		need := runtime.Sub(now)
		err = h.changeVisibility(ctx, url, msg, min(visibility, need))
	}

	return
}

func (h *Handle) process(ctx context.Context, msg *types.Message, handler message.Handler[any]) (err error) {
	peek := handler.Peek()
	if de := h.param.Decoder.Decode(msg.Body, peek); nil != de {
		err = de
	} else {
		extra := new(message.Extra)
		extra.Id = msg.MessageId
		extra.Handle = msg.ReceiptHandle
		extra.Attributes = msg.Attributes
		extra.Messages = msg.MessageAttributes
		err = handler.Process(kernel.New(ctx), peek, extra)
	}

	return
}

func (h *Handle) cleanup(ctx context.Context, url *string, msg *types.Message, err *error) {
	if nil == *err { // 消费成功，删除消息，不然会重复消费
		_ = h.delete(ctx, url, msg)
	} else if delay, ok := (*err).(*exception.Delay); ok { // 延迟消费，改变消息可见性，使其在指定的时间内再次被消费
		_ = h.changeVisibility(ctx, url, msg, delay.Duration())
	} else {
		_ = h.changeVisibility(ctx, url, msg, time.Duration(h.fibonacci(h.receiveCount(msg)))*time.Second)
	}
}

func (h *Handle) changeVisibility(
	ctx context.Context,
	url *string, msg *types.Message,
	visibility time.Duration,
) (err error) {
	cvi := new(sqs.ChangeMessageVisibilityInput)
	cvi.QueueUrl = url
	cvi.ReceiptHandle = msg.ReceiptHandle
	cvi.VisibilityTimeout = int32(visibility.Seconds())

	fields := gox.Fields[any]{
		field.New("id", msg.MessageId),
		field.New("next", time.Now().Add(visibility)),
	}
	if _, ve := h.param.Visibility(ctx, cvi); nil != err {
		err = ve
		h.logger.Warn("达到最大重试次数，改变消息可见性等待下一次消费", fields.Add(field.Error(ve))...)
	} else {
		h.logger.Info("达到最大重试次数，改变消息可见性等待下一次消费", fields...)
	}

	return
}

func (h *Handle) delete(ctx context.Context, url *string, msg *types.Message) (err error) {
	dmi := new(sqs.DeleteMessageInput)
	dmi.QueueUrl = url
	dmi.ReceiptHandle = msg.ReceiptHandle

	fields := gox.Fields[any]{
		field.New("id", msg.MessageId),
		field.New("receipt", msg.ReceiptHandle),
	}
	if _, de := h.param.Delete(ctx, dmi); nil != de {
		err = de
		h.logger.Info("删除消息出错", fields.Add(field.Error(de))...)
	} else {
		delete(h.counts, *msg.MessageId) // ! 强制删除缓存
		h.logger.Debug("删除消息成功", fields...)
	}

	return
}

func (h *Handle) renew(ctx context.Context, url *string, msg *types.Message) (err error) {
	smi := new(sqs.SendMessageInput)
	smi.QueueUrl = url
	smi.MessageAttributes = msg.MessageAttributes
	smi.MessageBody = msg.Body
	if out, se := h.receive.Send(ctx, smi); nil != se {
		err = se
	} else if "" != *out.MessageId {
		err = h.delete(ctx, url, msg)
	}

	return
}

func (h *Handle) getTime(msg *types.Message) (sent time.Time, run time.Time) {
	now := time.Now()
	runtime := *msg.MessageAttributes[constant.Runtime].StringValue
	if realtime, pe := time.Parse(time.RFC3339, runtime); nil == pe {
		run = realtime
	} else {
		run = now
	}
	if milliseconds, pe := strconv.ParseInt(msg.Attributes[constant.KeySentTimestamp], 10, 64); nil != pe {
		sent = time.UnixMilli(milliseconds)
	} else {
		sent = now
	}

	return
}

func (h *Handle) receiveCount(msg *types.Message) (count int) {
	id := *msg.MessageId
	if parsed, pe := strconv.Atoi(msg.Attributes[constant.KeyReceiveCount]); nil != pe {
		count = parsed
	} else if cached, ok := h.counts[id]; ok {
		count = cached
		h.counts[id] = count + 1
	} else {
		count = 1
		h.counts[id] = count
	}

	return
}

func (h *Handle) fibonacci(count int) (result int) {
	if count < 2 {
		result = 1
	} else {
		result = h.fibonacci(count-1) + h.fibonacci(count-2)
	}

	return
}
