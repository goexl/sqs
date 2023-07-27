package sqs

import (
	"github.com/goexl/sqs/internal/core"
)

var _ = New

// Client 客户端
// 使用 New 来创建
type Client = core.Client

func New() *core.Builder {
	return core.NewBuilder()
}
