package test_test

import (
	"context"

	"github.com/goexl/sqs"
)

type Handler struct{}

func (h *Handler) Peek() any {
	return new(User)
}

func (h *Handler) Process(_ context.Context, _ any, _ *sqs.Extra) (status sqs.Status, err error) {
	status = sqs.StatusSuccess

	return
}
