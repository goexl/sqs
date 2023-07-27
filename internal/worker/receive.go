package worker

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/goexl/gox"
	"github.com/goexl/sqs/internal/output"
	"github.com/goexl/sqs/internal/param"
)

type Receive struct {
	client *param.Client
	param  *param.Receive
}

func NewReceive(client *param.Client, param *param.Receive) *Receive {
	return &Receive{
		client: client,
		param:  param,
	}
}

func (r *Receive) Do(ctx context.Context) (out *output.Receive, err error) {
	if url, ue := r.param.Url(ctx, r.param.Label); nil != ue {
		err = ue
	} else {
		out, err = r.do(ctx, url)
	}

	return
}

func (r *Receive) do(ctx context.Context, url *string) (out *output.Receive, err error) {
	rmi := new(sqs.ReceiveMessageInput)
	rmi.QueueUrl = url
	rmi.AttributeNames = r.param.AttributeNames
	rmi.MaxNumberOfMessages = r.param.MaxNumberOfMessages
	rmi.MessageAttributeNames = r.param.MessageAttributeNames
	rmi.VisibilityTimeout = r.param.VisibilityTimeout
	rmi.WaitTimeSeconds = int32(gox.Ift(0 != r.param.Wait, r.param.Wait, r.client.Wait) / time.Second)

	if rsp, re := r.param.Receive(ctx, rmi); nil != re {
		err = re
	} else {
		out = rsp
	}

	return
}
