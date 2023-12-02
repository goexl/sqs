package internal

import (
	"github.com/goexl/sqs/internal/internal/constant"
)

type Base struct {
	Label string
	Queue string
	Url   string
}

func NewBase() *Base {
	return &Base{
		Label: constant.DefaultLabel,
	}
}
