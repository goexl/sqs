package credential

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/goexl/gox/http"
)

var _ aws.CredentialsProvider = (*Http)(nil)

type Http struct {
	method http.Method
	url    string
}

func NewHttp(method http.Method, url string) *Http {
	return &Http{
		method: method,
		url:    url,
	}
}

func (h *Http) Retrieve(_ context.Context) (credential aws.Credentials, err error) {
	credential = aws.Credentials{}

	return
}
