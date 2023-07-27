package param

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/goexl/sqs/internal/param/internal"
)

type Provider struct {
	Credentials aws.CredentialsProvider
}

func NewProvider() (provider *Provider) {
	internal.Once.Do(func() {
		provider = new(Provider)
	})

	return
}
