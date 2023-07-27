package core

import (
	"github.com/goexl/sqs/internal/builder"
)

func (c *Client) Receive() *builder.Receive {
	return builder.NewReceive(c.param, c.sqs.ReceiveMessage, c.url)
}

func (c *Client) Handle() *builder.Handle {
	return builder.NewHandle(c.param, c.sqs.ReceiveMessage, c.url, c.sqs.ChangeMessageVisibility, c.sqs.DeleteMessage)
}
