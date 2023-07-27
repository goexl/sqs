package context

import (
	"context"
	"time"
)

const (
	consumeContextKey = "sqs.consume"
	defaultDelay      = 5 * time.Second
)

type Consume struct {
	// 延迟时间
	delay time.Duration
}

func WithDelay(ctx context.Context, delay time.Duration) {
	if consume, ok := ctx.Value(consumeContextKey).(*Consume); ok {
		consume.delay = delay
	}
}

func WithConsume(ctx context.Context) context.Context {
	return context.WithValue(ctx, consumeContextKey, &Consume{
		delay: defaultDelay,
	})
}

func Delay(ctx context.Context) (delay time.Duration) {
	if consume, ok := ctx.Value(consumeContextKey).(*Consume); ok {
		delay = consume.delay
	} else {
		delay = defaultDelay
	}

	return
}
