package core

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"github.com/goexl/sqs/internal/internal"
)

func (c *Client) getAttributes(ctx context.Context, url *string) (attributes *internal.Attributes, err error) {
	gai := new(sqs.GetQueueAttributesInput)
	gai.QueueUrl = url
	gai.AttributeNames = []types.QueueAttributeName{
		types.QueueAttributeNameVisibilityTimeout,
		types.QueueAttributeNameMaximumMessageSize,
		types.QueueAttributeNameDelaySeconds,
		types.QueueAttributeNameMessageRetentionPeriod,
	}

	if cached, ok := c.attributes.Load(*url); ok {
		attributes = cached.(*internal.Attributes)
	} else if out, ge := c.sqs.GetQueueAttributes(ctx, gai); nil != ge {
		err = ge
	} else {
		attributes = internal.NewAttributes(out.Attributes)
		c.attributes.Store(*url, attributes)
	}

	return
}
