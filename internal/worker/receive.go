package worker

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/goexl/sqs/internal/output"
	"github.com/goexl/sqs/internal/param"
)

type Receive struct {
	param *param.Receive
}

func NewReceive(param *param.Receive) *Receive {
	return &Receive{
		param: param,
	}
}

func (r *Receive) Do(ctx context.Context) (out *output.Receive, err error) {
	if url, ue := r.param.Url(ctx, r.param.Base); nil != ue {
		err = ue
	} else {
		out, err = r.do(ctx, url)
	}

	return
}

func (r *Receive) do(ctx context.Context, url *string) (out *output.Receive, err error) {
	rmi := new(sqs.ReceiveMessageInput)
	rmi.QueueUrl = url
	rmi.AttributeNames = r.param.Names
	rmi.MaxNumberOfMessages = r.param.Number
	rmi.MessageAttributeNames = r.param.Attributes
	rmi.VisibilityTimeout = r.param.Visibility
	rmi.WaitTimeSeconds = r.param.WaitTimeSeconds()

	if rsp, re := r.param.Receive(ctx, rmi); nil != re {
		err = re
	} else {
		out = rsp
	}

	return
}
