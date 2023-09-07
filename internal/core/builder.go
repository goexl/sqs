package core

import (
	"time"

	"github.com/goexl/simaqian"
	"github.com/goexl/sqs/internal/param"
)

type Builder struct {
	param *param.Client
}

func NewBuilder() *Builder {
	return &Builder{
		param: param.NewClient(),
	}
}

func (b *Builder) Logger(logger simaqian.Logger) (client *Builder) {
	b.param.Logger = logger
	client = b

	return
}

func (b *Builder) Region(region string) (client *Builder) {
	b.param.Region = region
	client = b

	return
}

func (b *Builder) Wait(wait time.Duration) (client *Builder) {
	b.param.Wait = wait
	client = b

	return
}

func (b *Builder) Queue(label string, name string) (client *Builder) {
	b.param.Queues[label] = &name
	client = b

	return
}

func (b *Builder) Credential() *Credential {
	return NewCredential(b, b.param)
}

func (b *Builder) Build() *Client {
	return NewClient(b.param)
}
