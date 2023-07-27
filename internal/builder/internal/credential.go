package internal

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/goexl/gox/http"
	"github.com/goexl/sqs/internal/builder"
	"github.com/goexl/sqs/internal/param"
	"github.com/goexl/sqs/internal/provider/credential"
)

type Credential struct {
	client   *builder.Client
	param    *param.Client
	provider aws.CredentialsProvider
}

func NewCredential(client *builder.Client, param *param.Client) *Credential {
	return &Credential{
		client: client,
		param:  param,
	}
}

func (c *Credential) Default(access string, secret string) (cdl *Credential) {
	c.provider = credential.NewDefault(access, secret)
	cdl = c

	return
}

func (c *Credential) Http(method http.Method, url string) (cdl *Credential) {
	c.provider = credential.NewHttp(method, url)
	cdl = c

	return
}

func (c *Credential) Build() (client *builder.Client) {
	c.param.Provider = c.provider
	client = c.client

	return
}
