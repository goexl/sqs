package test_test

import (
	"github.com/goexl/sqs"
	"github.com/goexl/sqs/internal/kernel"
)

type Handler struct{}

func (h *Handler) Peek() any {
	return new(User)
}

func (h *Handler) Process(_ *kernel.Context, _ any, _ *sqs.Extra) (err error) {
	return
}
