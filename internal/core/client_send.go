package core

import (
	"github.com/goexl/sqs/internal/internal/builder"
)

func (c *Client) Send() *builder.Send {
	return builder.NewSend(c.sqs.SendMessage, c.url)
}
