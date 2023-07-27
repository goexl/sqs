package builder

import (
	"time"

	"github.com/goexl/sqs/internal/callback"
	"github.com/goexl/sqs/internal/param"
	"github.com/goexl/sqs/internal/transcoder"
	"github.com/goexl/sqs/internal/worker"
)

type Handle struct {
	*Receive

	param  *param.Handle
	change callback.ChangeMessageVisibility
	delete callback.DeleteMessage
}

func NewHandle(
	client *param.Client,
	receive callback.ReceiveMessage,
	url callback.Url,
	change callback.ChangeMessageVisibility,
	delete callback.DeleteMessage,
) *Handle {
	return &Handle{
		Receive: NewReceive(client, receive, url),

		param:  param.NewHandle(),
		change: change,
		delete: delete,
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
	return worker.NewHandle(h.Receive.param, h.param)
}
