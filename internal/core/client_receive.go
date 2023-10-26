package core

import (
	"github.com/goexl/sqs/internal/builder"
)

func (c *Client) Receive() *builder.Receive {
	return builder.NewReceive(c.params, c.sqs.SendMessage, c.sqs.ReceiveMessage, c.url, c.getAttributes)
}

func (c *Client) Handle() *builder.Handle {
	return builder.NewHandle(
		c.params,
		c.sqs.SendMessage, c.sqs.ReceiveMessage,
		c.url, c.sqs.ChangeMessageVisibility, c.sqs.DeleteMessage, c.getAttributes,
	)
}
