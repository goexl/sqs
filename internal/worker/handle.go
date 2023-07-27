package worker

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
	"github.com/goexl/simaqian"
	"github.com/goexl/sqs/internal/context"
	"github.com/goexl/sqs/internal/message"
	"github.com/goexl/sqs/internal/param"
)

type Handle struct {
	simaqian.Logger

	receive *param.Receive
	param   *param.Handle
}

func NewHandle(receive *param.Receive, param *param.Handle) *Handle {
	return &Handle{
		receive: receive,
		param:   param,
	}
}

func (h *Handle) Start(ctx context.Context, handler message.Handler[any]) (err error) {
	if url, ue := h.receive.Url(ctx, h.receive.Label); nil != ue {
		err = ue
	} else {
		err = h.do(ctx, url, handler)
	}

	return

}

func (h *Handle) do(ctx context.Context, url *string, handler message.Handler[any]) (err error) {
	rmi := new(sqs.ReceiveMessageInput)
	rmi.QueueUrl = url
	rmi.AttributeNames = h.receive.Names
	rmi.MaxNumberOfMessages = h.receive.Number
	rmi.MessageAttributeNames = h.receive.Attributes
	rmi.VisibilityTimeout = h.receive.Visibility
	rmi.WaitTimeSeconds = h.receive.WaitTimeSeconds()

	for {
		if rsp, re := h.receive.Receive(ctx, rmi); nil != err {
			return
		} else if 1 != len(rsp.Messages) {
			h.Warn("收取消息出错", field.New("url", url), field.Error(re))
		} else { // 并行消费，加快消费速度
			for _, msg := range rsp.Messages {
				go h.handle(ctx, url, &msg, handler)
			}
		}
	}
}

func (h *Handle) handle(ctx context.Context, url *string, msg *types.Message, handler message.Handler[any]) {
	var err error
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
		_ = h.visibility(ctx, url, msg, h.param.Interval)
	} else {
		h.status(ctx, url, msg, status)
	}
}

func (h *Handle) status(ctx context.Context, url *string, msg *types.Message, status *message.Status) {
	switch *status {
	case message.StatusSuccess: // 消费成功，删除消息，不然会重复消费
		_ = h.delete(ctx, url, msg)
	case message.StatusLater: // 延迟消费，改变消息可见性，使其在指定的时间内再次被消费
		_ = h.visibility(ctx, url, msg, context.Delay(ctx))
	case message.StatusUnknown: // 默认状态，改变消息可见性，使前可以立即被消费
		_ = h.visibility(ctx, url, msg, time.Second)
	}

	return
}

func (h *Handle) visibility(ctx context.Context, url *string, msg *types.Message, timeout time.Duration) (err error) {
	cvi := new(sqs.ChangeMessageVisibilityInput)
	cvi.QueueUrl = url
	cvi.ReceiptHandle = msg.ReceiptHandle
	cvi.VisibilityTimeout = int32(timeout / time.Second)

	fields := gox.Fields[any]{
		field.New("id", msg.MessageId),
		field.New("next", time.Now().Add(timeout)),
	}
	if _, ve := h.param.Visibility(ctx, cvi); nil != err {
		err = ve
		h.Info("达到最大重试次数，改变消息可见性等待下一次消费", fields.Add(field.Error(ve))...)
	} else {
		h.Debug("达到最大重试次数，改变消息可见性等待下一次消费", fields...)
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
		h.Info("删除消息出错", fields.Add(field.Error(de))...)
	} else {
		h.Debug("删除消息成功", fields...)
	}

	return
}
