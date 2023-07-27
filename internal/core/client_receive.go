package core

import (
	"github.com/goexl/sqs/internal/builder"
)

func (c *Client) Receive() *builder.Receive {
	return builder.NewReceive(c.param, c.client.ReceiveMessage, c.Url)
}

func (c *Client) Handle() *builder.Handle {
	return builder.NewHandle(c.client.ReceiveMessage, c.Url, c.client.ChangeMessageVisibility, c.client.DeleteMessage)
}
