package core

import (
	builder2 "github.com/goexl/sqs/internal/internal/builder"
)

func (c *Client) Receive() *builder2.Receive {
	return builder2.NewReceive(c.params, c.sqs.SendMessage, c.sqs.ReceiveMessage, c.url, c.getAttributes)
}

func (c *Client) Handle() *builder2.Handle {
	return builder2.NewHandle(
		c.params,
		c.sqs.SendMessage, c.sqs.ReceiveMessage,
		c.url, c.sqs.ChangeMessageVisibility, c.sqs.DeleteMessage, c.getAttributes,
	)
}
