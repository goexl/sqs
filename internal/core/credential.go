package core

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/goexl/gox/http"
	"github.com/goexl/sqs/internal/param"
	"github.com/goexl/sqs/internal/provider/credential"
)

type Credential struct {
	client   *Builder
	param    *param.Client
	provider aws.CredentialsProvider
}

func NewCredential(client *Builder, param *param.Client) *Credential {
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

func (c *Credential) Build() (client *Builder) {
	c.param.Provider.Credentials = c.provider
	client = c.client

	return
}
