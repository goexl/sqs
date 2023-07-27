package param

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"github.com/goexl/gox"
	"github.com/goexl/sqs/internal/callback"
	"github.com/goexl/sqs/internal/internal"
)

type Receive struct {
	*internal.Base

	Wait       time.Duration
	Visibility int32
	Number     int32
	Names      []types.QueueAttributeName
	Attributes []string

	Receive callback.ReceiveMessage
	Url     callback.Url

	client *Client
}

func NewReceive(client *Client, receive callback.ReceiveMessage, url callback.Url) *Receive {
	return &Receive{
		Base: internal.NewBase(),

		Wait:   15 * time.Second,
		Number: 1,

		Receive: receive,
		Url:     url,

		client: client,
	}
}

func (r *Receive) WaitTimeSeconds() int32 {
	return int32(gox.Ift(0 != r.Wait, r.Wait, r.client.Wait) / time.Second)
}
