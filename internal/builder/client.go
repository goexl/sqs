package builder

import (
	"time"

	"github.com/goexl/sqs/internal/builder/internal"
	"github.com/goexl/sqs/internal/core"
	"github.com/goexl/sqs/internal/param"
)

type Client struct {
	param *param.Client
}

func NewClient() *Client {
	return &Client{
		param: param.NewClient(),
	}
}

func (c *Client) Region(region string) (client *Client) {
	c.param.Region = region
	client = c

	return
}

func (c *Client) Wait(wait time.Duration) (client *Client) {
	c.param.Wait = wait
	client = c

	return
}

func (c *Client) Credential() *internal.Credential {
	return internal.NewCredential(c, c.param)
}

func (c *Client) Build() *core.Client {
	return core.NewClient(c.param)
}
