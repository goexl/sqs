package builder

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/goexl/sqs/internal/param"
)

type Provider struct {
	param *param.Provider
}

func NewProvider() *Provider {
	return &Provider{
		param: param.NewProvider(),
	}
}

func (p *Provider) Credentials(credentials aws.CredentialsProvider) (provider *Provider) {
	p.param.Credentials = credentials
	provider = p

	return
}
