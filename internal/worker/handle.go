package worker

import (
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
	"github.com/goexl/log"
	"github.com/goexl/sqs/internal/constant"
	"github.com/goexl/sqs/internal/context"
	"github.com/goexl/sqs/internal/message"
	"github.com/goexl/sqs/internal/param"
)

type Handle struct {
	logger  log.Logger
	receive *param.Receive
	param   *param.Handle
}

func NewHandle(logger log.Logger, receive *param.Receive, param *param.Handle) *Handle {
	return &Handle{
		logger:  logger,
		receive: receive,
		param:   param,
	}
}

func (h *Handle) Start(ctx context.Context, handler message.Handler[any]) (err error) {
	if url, ue := h.receive.Url(ctx, h.receive.Base); nil != ue {
		err = ue
	} else {
		err = h.do(ctx, url, handler)
	}

	return

}

func (h *Handle) do(ctx context.Context, url *string, handler message.Handler[any]) (err error) {
	for {
		if out, re := h.receive.Do(ctx, url); nil != re {
			h.logger.Warn("收取消息出错", field.New("url", url), field.Error(re))
		} else { // 并行消费，加快消费速度
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
	status := message.StatusUnknown
	defer h.cleanup(ctx, url, msg, &status, &err)

	for times := 0; times < h.param.Times; times++ {
		status, err = h.process(ctx, msg, handler)
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

func (h *Handle) process(
	ctx context.Context,
	msg *types.Message,
	handler message.Handler[any],
) (status message.Status, err error) {
	peek := handler.Peek()
	if de := h.param.Decoder.Decode(msg.Body, peek); nil != de {
		err = de
	} else {
		extra := new(message.Extra)
		extra.Id = msg.MessageId
		extra.Handle = msg.ReceiptHandle
		extra.Attributes = msg.Attributes
		extra.Messages = msg.MessageAttributes

		status, err = handler.Process(context.WithConsume(ctx), peek, extra)
	}

	return
}

func (h *Handle) cleanup(ctx context.Context, url *string, msg *types.Message, status *message.Status, err *error) {
	if nil != *err {
		_ = h.changeVisibility(ctx, url, msg, h.param.Interval)
	} else {
		h.status(ctx, url, msg, status)
	}
}

func (h *Handle) status(ctx context.Context, url *string, msg *types.Message, status *message.Status) {
	switch *status {
	case message.StatusSuccess: // 消费成功，删除消息，不然会重复消费
		_ = h.delete(ctx, url, msg)
	case message.StatusLater: // 延迟消费，改变消息可见性，使其在指定的时间内再次被消费
		_ = h.changeVisibility(ctx, url, msg, context.Delay(ctx))
	case message.StatusUnknown: // 默认状态，改变消息可见性，使前可以立即被消费
		_ = h.changeVisibility(ctx, url, msg, time.Second)
	}

	return
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
		h.logger.Info("达到最大重试次数，改变消息可见性等待下一次消费", fields.Add(field.Error(ve))...)
	} else {
		h.logger.Debug("达到最大重试次数，改变消息可见性等待下一次消费", fields...)
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
