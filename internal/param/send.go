package param

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"github.com/goexl/sqs/internal/callback"
	"github.com/goexl/sqs/internal/internal"
	"github.com/goexl/sqs/internal/transcoder"
)

type Send struct {
	*internal.Base
	*Provider

	Data       any
	Delay      time.Duration
	Runtime    *time.Time
	Attributes map[string]types.MessageAttributeValue
	Systems    map[string]types.MessageSystemAttributeValue

	Encoder       transcoder.Encoder
	Send          callback.SendMessage
	Url           callback.Url
	GetAttributes callback.GetAttributes
}

func NewSend(send callback.SendMessage, url callback.Url) *Send {
	return &Send{
		Base:     internal.NewBase(),
		Provider: NewProvider(),

		Attributes: make(map[string]types.MessageAttributeValue),
		Systems:    make(map[string]types.MessageSystemAttributeValue),

		Encoder: transcoder.NewJson(),
		Send:    send,
		Url:     url,
	}
}
