package param

import (
	"time"

	"github.com/goexl/sqs/internal/callback"
	"github.com/goexl/sqs/internal/internal"
	"github.com/goexl/sqs/internal/transcoder"
)

type Handle struct {
	*internal.Base
	*Provider

	Times    int
	Interval time.Duration

	Decoder       transcoder.Decoder
	Visibility    callback.ChangeMessageVisibility
	Delete        callback.DeleteMessage
	GetAttributes callback.GetAttributes
}

func NewHandle(visibility callback.ChangeMessageVisibility, delete callback.DeleteMessage) *Handle {
	return &Handle{
		Base:     internal.NewBase(),
		Provider: NewProvider(),

		Times:    3,
		Interval: 3 * time.Second,

		Decoder:    transcoder.NewJson(),
		Visibility: visibility,
		Delete:     delete,
	}
}
