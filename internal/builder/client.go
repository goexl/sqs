package builder

import (
	"time"

	"github.com/goexl/http"
	"github.com/goexl/log"
	"github.com/goexl/sqs/internal/core"
	"github.com/goexl/sqs/internal/param"
)

type Client struct {
	param *param.Client
}

func NewBuilder() *Client {
	return &Client{
		param: param.NewClient(),
	}
}

func (c *Client) Logger(logger log.Logger) (client *Client) {
	c.param.Logger = logger
	client = c

	return
}

func (c *Client) Http(http *http.Client) (client *Client) {
	c.param.Http = http
	client = c

	return
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

func (c *Client) Queue(label string, name string) (client *Client) {
	c.param.Queues[label] = &name
	client = c

	return
}

func (c *Client) Credential() *Credential {
	return NewCredential(c, c.param)
}

func (c *Client) Build() *core.Client {
	return core.NewClient(c.param)
}
