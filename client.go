package sqs

import (
	"github.com/goexl/sqs/internal/builder"
	"github.com/goexl/sqs/internal/core"
)

var _ = New

// Client 客户端
// 使用 New 来创建
type Client = core.Client

func New() *builder.Client {
	return builder.NewBuilder()
}
