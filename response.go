package sqs

import (
	"github.com/goexl/sqs/internal/output"
)

type (
	// SendResponse 发送响应
	SendResponse = output.Send
	// ReceiveResponse 接收响应
	ReceiveResponse = output.Receive
)
