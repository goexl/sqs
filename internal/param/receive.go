package param

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"github.com/goexl/sqs/internal/callback"
	"github.com/goexl/sqs/internal/param/internal"
)

type Receive struct {
	internal.Base

	Wait                  time.Duration
	VisibilityTimeout     int32
	MaxNumberOfMessages   int32
	AttributeNames        []types.QueueAttributeName
	MessageAttributeNames []string

	Receive callback.ReceiveMessage
	Url     callback.Url
}

func NewReceive(receive callback.ReceiveMessage, url callback.Url) *Receive {
	return &Receive{
		Wait:                15 * time.Second,
		MaxNumberOfMessages: 1,

		Receive: receive,
		Url:     url,
	}
}
