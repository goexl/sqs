package kernel

import (
	"context"
	"time"

	"github.com/goexl/sqs/internal/internal/exception"
)

type Context struct {
	context.Context
}

func New(parent context.Context) *Context {
	return &Context{
		Context: parent,
	}
}

func (c *Context) Delay(delay time.Duration) error {
	return exception.NewDelay(&delay)
}
