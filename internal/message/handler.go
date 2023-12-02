package message

import (
	"github.com/goexl/sqs/internal/kernel"
)

type Handler[T any] interface {
	// Peek 取出一个新的消息
	Peek() T

	// Process 处理消息
	Process(ctx *kernel.Context, msg T, extra *Extra) error
}
