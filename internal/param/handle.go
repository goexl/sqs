package param

import (
	"time"

	"github.com/goexl/sqs/internal/callback"
	"github.com/goexl/sqs/internal/transcoder"
)

type Handle struct {
	Times    int
	Interval time.Duration

	Decoder    transcoder.Decoder
	Visibility callback.ChangeMessageVisibility
	Delete     callback.DeleteMessage
}

func NewHandle() *Handle {
	return &Handle{
		Times:    3,
		Interval: 3 * time.Second,
	}
}
