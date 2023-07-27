package builder

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"github.com/goexl/sqs/internal/callback"
	"github.com/goexl/sqs/internal/param"
	"github.com/goexl/sqs/internal/worker"
)

type Receive struct {
	param   *param.Receive
	receive callback.ReceiveMessage
	url     callback.Url
}

func NewReceive(client *param.Client, receive callback.ReceiveMessage, url callback.Url) *Receive {
	return &Receive{
		param:   param.NewReceive(client, receive, url),
		receive: receive,
		url:     url,
	}
}

func (r *Receive) Names(names ...types.QueueAttributeName) (receive *Receive) {
	r.param.Names = names
	receive = r

	return
}

func (r *Receive) Label(label string) (receive *Receive) {
	r.param.Label = label
	receive = r

	return
}

func (r *Receive) Number(number int32) (receive *Receive) {
	r.param.Number = number
	receive = r

	return
}

func (r *Receive) Attributes(names ...string) (receive *Receive) {
	r.param.Attributes = names
	receive = r

	return
}

func (r *Receive) Visibility(timeout time.Duration) (receive *Receive) {
	r.param.Visibility = int32(timeout / time.Second)
	receive = r

	return
}

func (r *Receive) Wait(wait time.Duration) (receive *Receive) {
	r.param.Wait = wait
	receive = r

	return
}

func (r *Receive) Build() *worker.Receive {
	return worker.NewReceive(r.param)
}
