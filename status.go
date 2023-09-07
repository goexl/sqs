package sqs

import (
	"github.com/goexl/sqs/internal/message"
)

// Status 状态
type Status = message.Status

const (
	StatusSuccess = message.StatusSuccess
	StatusLater   = message.StatusLater
)

var (
	_ = StatusSuccess
	_ = StatusLater
)
