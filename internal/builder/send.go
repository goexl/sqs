package builder

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"github.com/goexl/sqs/internal/callback"
	"github.com/goexl/sqs/internal/param"
	"github.com/goexl/sqs/internal/transcoder"
	"github.com/goexl/sqs/internal/worker"
)

type Send struct {
	*Base

	param        *param.Send
	defaultDelay time.Duration
}

func NewSend(send callback.SendMessage, url callback.Url) *Send {
	return &Send{
		Base: NewBase(),

		param:        param.NewSend(send, url),
		defaultDelay: 15 * time.Minute,
	}
}

func (s *Send) Delay(delay time.Duration) (send *Send) {
	if delay <= s.defaultDelay {
		s.param.Delay = delay
	} else {
		diff := delay - s.defaultDelay
		runtime := time.Now().Add(diff)
		s.param.Runtime = &runtime
	}
	send = s

	return
}

func (s *Send) Fix(set time.Time) (send *Send) {
	diff := set.Sub(time.Now())
	if diff <= s.defaultDelay {
		s.param.Delay = diff
	} else {
		s.param.Delay = s.defaultDelay
		runtime := set.Add(-s.defaultDelay)
		s.param.Runtime = &runtime
	}
	send = s

	return
}

func (s *Send) Label(label string) (send *Send) {
	s.param.Label = label
	send = s

	return
}

func (s *Send) Data(data any) (send *Send) {
	s.param.Data = data
	send = s

	return
}

func (s *Send) Encoder(encoder transcoder.Encoder) (send *Send) {
	s.param.Encoder = encoder
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
