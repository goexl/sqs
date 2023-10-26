package param

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"github.com/goexl/gox"
	"github.com/goexl/sqs/internal/callback"
	"github.com/goexl/sqs/internal/constant"
	"github.com/goexl/sqs/internal/context"
	"github.com/goexl/sqs/internal/internal"
)

type Receive struct {
	*internal.Base
	*Provider

	Wait       time.Duration
	Visibility time.Duration
	Number     int32
	Names      []types.QueueAttributeName
	Attributes []string

	Send          callback.SendMessage
	Url           callback.Url
	GetAttributes callback.GetAttributes

	client  *Client
	receive callback.ReceiveMessage
}

func NewReceive(
	client *Client,
	send callback.SendMessage, receive callback.ReceiveMessage,
	url callback.Url, attributes callback.GetAttributes,
) *Receive {
	return &Receive{
		Base:     internal.NewBase(),
		Provider: NewProvider(),

		Wait:   15 * time.Second,
		Number: 1,

		Send:          send,
		Url:           url,
		GetAttributes: attributes,

		client:  client,
		receive: receive,
	}
}

func (r *Receive) Do(ctx context.Context, url *string) (*sqs.ReceiveMessageOutput, error) {
	input := new(sqs.ReceiveMessageInput)
	input.QueueUrl = url
	input.AttributeNames = append(r.Names, constant.KeySentTimestamp)
	input.MaxNumberOfMessages = r.Number
	input.MessageAttributeNames = append(r.Attributes, constant.Runtime)
	input.VisibilityTimeout = int32(r.Visibility.Seconds())
	input.WaitTimeSeconds = int32(gox.Ift(0 != r.Wait, r.Wait, r.client.Wait).Seconds())

	return r.receive(ctx, input)
}

func (r *Receive) Exited() bool {
	return r.client.Exit
}

func (r *Receive) Cancel(cancel context.CancelFunc) {
	r.client.Cancel = cancel
}
