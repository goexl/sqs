package worker

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/goexl/sqs/internal/output"
	"github.com/goexl/sqs/internal/param"
)

type Send struct {
	param *param.Send
}

func NewSend(param *param.Send) *Send {
	return &Send{
		param: param,
	}
}

func (s *Send) Do(ctx context.Context) (out *output.Send, err error) {
	if url, ue := s.param.Url(ctx, s.param.Label); nil != ue {
		err = ue
	} else {
		out, err = s.do(ctx, url)
	}

	return
}

func (s *Send) do(ctx context.Context, url *string) (out *output.Send, err error) {
	smi := new(sqs.SendMessageInput)
	smi.QueueUrl = url
	smi.DelaySeconds = int32(s.param.Delay / time.Second)
	smi.MessageAttributes = s.param.Attributes
	smi.MessageSystemAttributes = s.param.Systems

	if rsp, se := s.param.Send(ctx, smi); nil != se {
		err = se
	} else {
		out = rsp
	}

	return
}
