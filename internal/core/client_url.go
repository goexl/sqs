package core

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
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
	gqu.QueueName = c.queues[label]

	if rsp, ge := c.sqs.GetQueueUrl(ctx, gqu); nil != ge {
		err = ge
	} else {
		url = rsp.QueueUrl
		c.urls.Store(label, url)
	}

	return
}
