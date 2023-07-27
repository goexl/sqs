package core

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/goexl/exc"
	"github.com/goexl/gox/field"
)

func (c *Client) Url(ctx context.Context, label string) (url *string, err error) {
	if cache, ok := c.urls.Load(label); ok {
		url = cache.(*string)
	} else {
		url, err = c.url(ctx, label)
	}

	return
}

func (c *Client) url(ctx context.Context, label string) (url *string, err error) {
	gqu := new(sqs.GetQueueUrlInput)
	gqu.QueueName = c.param.Queues[label]

	if "" == *gqu.QueueName {
		err = exc.NewField("必须指定除名名称", field.New("label", label))
	} else if rsp, ge := c.sqs.GetQueueUrl(ctx, gqu); nil != ge {
		err = ge
	} else {
		url = rsp.QueueUrl
		c.urls.Store(label, url)
	}

	return
}
