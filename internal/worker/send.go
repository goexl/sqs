package worker

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"github.com/goexl/sqs/internal/constant"
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
	if url, ue := s.param.Url(ctx, s.param.Base); nil != ue {
		err = ue
	} else {
		out, err = s.do(ctx, url)
	}

	return
}

func (s *Send) do(ctx context.Context, url *string) (out *output.Send, err error) {
	smi := new(sqs.SendMessageInput)
	smi.QueueUrl = url
	smi.MessageAttributes = s.param.Attributes
	smi.MessageSystemAttributes = s.param.Systems

	if 0 != s.param.Delay {
		smi.DelaySeconds = int32(s.param.Delay.Seconds())
	}
	if nil != s.param.Runtime {
		smi.MessageAttributes[constant.Runtime] = types.MessageAttributeValue{
			DataType:    aws.String(constant.DataTypeString),
			StringValue: aws.String(s.param.Runtime.Format(time.RFC3339)),
		}
	}

	if encoded, ee := s.param.Encoder.Encode(s.param.Data); nil != ee {
		err = ee
	} else {
		smi.MessageBody = encoded
		out, err = s.send(ctx, smi)
	}

	return
}

func (s *Send) send(ctx context.Context, smi *sqs.SendMessageInput) (out *output.Send, err error) {
	if rsp, se := s.param.Send(ctx, smi); nil != se {
		err = se
	} else {
		out = rsp
	}

	return
}
