package context

import (
	"context"
)

func WithCancel(ctx Context) (Context, context.CancelFunc) {
	return context.WithCancel(ctx)
}
