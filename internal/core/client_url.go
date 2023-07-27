package core

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/goexl/exc"
	"github.com/goexl/gox/field"
	"github.com/goexl/sqs/internal/internal"
)

func (c *Client) Url(ctx context.Context, label string) (url *string, err error) {
	if cache, ok := c.urls.Load(label); ok {
		url = cache.(*string)
	} else {
		// url, err = c.url(ctx, label)
	}

	return
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
		gqu.QueueName = c.param.Queues[base.Label]
	}

	if "" == *gqu.QueueName {
		err = exc.NewField("必须指定队列名称", field.New("label", base.Label))
	} else if rsp, ge := c.sqs.GetQueueUrl(ctx, gqu); nil != ge {
		err = ge
	} else {
		url = rsp.QueueUrl
		c.urls.Store(base.Label, url)
	}

	return
}
