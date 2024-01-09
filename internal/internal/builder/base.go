package builder

import (
	"github.com/goexl/sqs/internal/internal"
)

type Base struct {
	*Provider

	param *internal.Base
}

func NewBase() *Base {
	return &Base{
		Provider: NewProvider(),

		param: internal.NewBase(),
	}
}

func (b *Base) Url(url string) (base *Base) {
	b.param.Url = url
	base = b

	return
}

func (b *Base) Queue(queue string) (base *Base) {
	b.param.Queue = queue
	base = b

	return
}

func (b *Base) Label(label string) (base *Base) {
	b.param.Label = label
	base = b

	return
}
