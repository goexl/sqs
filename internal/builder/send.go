package builder

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"github.com/goexl/sqs/internal/callback"
	"github.com/goexl/sqs/internal/param"
	"github.com/goexl/sqs/internal/worker"
)

type Send struct {
	param *param.Send
	send  callback.SendMessage
	url   callback.Url
}

func NewSend(send callback.SendMessage, url callback.Url) *Send {
	return &Send{
		param: param.NewSend(send, url),
		send:  send,
		url:   url,
	}
}

func (s *Send) Delay(delay time.Duration) (send *Send) {
	s.param.Delay = delay
	send = s

	return
}

func (s *Send) Label(label string) (send *Send) {
	s.param.Label = label
	send = s

	return
}

func (s *Send) Attribute(key string, value types.MessageAttributeValue) (send *Send) {
	s.param.Attributes[key] = value
	send = s

	return
}

func (s *Send) System(key string, value types.MessageSystemAttributeValue) (send *Send) {
	s.param.Systems[key] = value
	send = s

	return
}

func (s *Send) Build() *worker.Send {
	return worker.NewSend(s.param)
}
