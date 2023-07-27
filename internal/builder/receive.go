package builder

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"github.com/goexl/sqs/internal/callback"
	"github.com/goexl/sqs/internal/param"
	"github.com/goexl/sqs/internal/worker"
)

type Receive struct {
	client  *param.Client
	param   *param.Receive
	receive callback.ReceiveMessage
	url     callback.Url
}

func NewReceive(client *param.Client, receive callback.ReceiveMessage, url callback.Url) *Receive {
	return &Receive{
		client:  client,
		param:   param.NewReceive(receive, url),
		receive: receive,
		url:     url,
	}
}

func (r *Receive) AttributeNames(names ...types.QueueAttributeName) (receive *Receive) {
	r.param.AttributeNames = names
	receive = r

	return
}

func (r *Receive) Label(label string) (receive *Receive) {
	r.param.Label = label
	receive = r

	return
}

func (r *Receive) MaxNumberOfMessages(max int32) (receive *Receive) {
	r.param.MaxNumberOfMessages = max
	receive = r

	return
}

func (r *Receive) MessageAttributeNames(names ...string) (receive *Receive) {
	r.param.MessageAttributeNames = names
	receive = r

	return
}

func (r *Receive) VisibilityTimeout(timeout time.Duration) (receive *Receive) {
	r.param.VisibilityTimeout = int32(timeout / time.Second)
	receive = r

	return
}

func (r *Receive) Wait(wait time.Duration) (receive *Receive) {
	r.param.Wait = wait
	receive = r

	return
}

func (r *Receive) Build() *worker.Receive {
	return worker.NewReceive(r.client, r.param)
}
