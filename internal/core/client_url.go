package core

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/goexl/exc"
	"github.com/goexl/gox/field"
	"github.com/goexl/sqs/internal/builder"
	"github.com/goexl/sqs/internal/internal"
)

func (c *Client) Url() *builder.Url {
	return builder.NewUrl(c.url)
}

func (c *Client) url(ctx context.Context, base *internal.Base) (url *string, err error) {
	if "" != base.Url {
		url = &base.Url
	} else {
		url, err = c.query(ctx, base)
	}

	return
}

func (c *Client) query(ctx context.Context, base *internal.Base) (url *string, err error) {
	gqu := new(sqs.GetQueueUrlInput)
	if "" != base.Queue {
		gqu.QueueName = &base.Queue
	} else {
		gqu.QueueName = c.params.Queues[base.Label]
	}

	label := base.Label
	if "" == *gqu.QueueName {
		err = exc.NewField("必须指定队列名称", field.New("label", base.Label))
	} else if cached, ok := c.urls.Load(label); ok {
		url = cached.(*string)
	} else if rsp, gue := c.sqs.GetQueueUrl(ctx, gqu); nil != gue {
		err = gue
	} else {
		url = rsp.QueueUrl
		c.urls.Store(label, url)
	}

	return
}
