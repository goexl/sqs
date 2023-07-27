package builder

import (
	"time"

	"github.com/goexl/sqs/internal/callback"
	"github.com/goexl/sqs/internal/param"
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

func (h *Handle) RetryMax(max int) (handle *Handle) {
	h.param.MaxRetry = max
	handle = h

	return
}

func (h *Handle) RetryDuration(duration time.Duration) (handle *Handle) {
	h.param.RetryDuration = duration
	handle = h

	return
}

func (h *Handle) Build() *worker.Handle {
	return worker.NewHandle(h.Receive.param, h.param)
}
