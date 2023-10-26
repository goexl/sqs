package builder

import (
	"time"

	"github.com/goexl/log"
	"github.com/goexl/sqs/internal/callback"
	"github.com/goexl/sqs/internal/param"
	"github.com/goexl/sqs/internal/transcoder"
	"github.com/goexl/sqs/internal/worker"
)

type Handle struct {
	*Receive

	param             *param.Handle
	messageVisibility callback.ChangeMessageVisibility
	delete            callback.DeleteMessage
	logger            log.Logger
}

func NewHandle(
	client *param.Client,
	send callback.SendMessage, receive callback.ReceiveMessage,
	url callback.Url, visibility callback.ChangeMessageVisibility,
	delete callback.DeleteMessage, attributes callback.GetAttributes,
) *Handle {
	return &Handle{
		Receive: NewReceive(client, send, receive, url, attributes),

		param:             param.NewHandle(visibility, delete),
		messageVisibility: visibility,
		delete:            delete,
		logger:            client.Logger,
	}
}

func (h *Handle) Times(max int) (handle *Handle) {
	h.param.Times = max
	handle = h

	return
}

func (h *Handle) Interval(duration time.Duration) (handle *Handle) {
	h.param.Interval = duration
	handle = h

	return
}

func (h *Handle) Decoder(decoder transcoder.Decoder) (handle *Handle) {
	h.param.Decoder = decoder
	handle = h

	return
}

func (h *Handle) Build() *worker.Handle {
	return worker.NewHandle(h.logger, h.Receive.param, h.param)
}
