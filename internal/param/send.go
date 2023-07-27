package param

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"github.com/goexl/sqs/internal/callback"
	"github.com/goexl/sqs/internal/param/internal"
	"github.com/goexl/sqs/internal/transcoder"
)

type Send struct {
	internal.Base

	Delay      time.Duration
	Attributes map[string]types.MessageAttributeValue
	Systems    map[string]types.MessageSystemAttributeValue

	Encoder transcoder.Encoder
	Send    callback.SendMessage
	Url     callback.Url
}

func NewSend(send callback.SendMessage, url callback.Url) *Send {
	return &Send{
		Attributes: make(map[string]types.MessageAttributeValue),
		Systems:    make(map[string]types.MessageSystemAttributeValue),

		Encoder: transcoder.NewJson(),
		Send:    send,
		Url:     url,
	}
}
