package param

import (
	"time"

	"github.com/goexl/sqs/internal/callback"
	"github.com/goexl/sqs/internal/transcoder"
)

type Handle struct {
	MaxRetry      int
	RetryDuration time.Duration

	Decoder    transcoder.Decoder
	Visibility callback.ChangeMessageVisibility
	Delete     callback.DeleteMessage
}

func NewHandle() *Handle {
	return &Handle{
		MaxRetry:      3,
		RetryDuration: 3 * time.Second,
	}
}
