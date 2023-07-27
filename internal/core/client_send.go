package core

import (
	"github.com/goexl/sqs/internal/builder"
)

func (c *Client) Send() *builder.Send {
	return builder.NewSend(c.sqs.SendMessage, c.Url)
}
