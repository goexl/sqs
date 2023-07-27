package credential

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
)

var _ aws.CredentialsProvider = (*Default)(nil)

type Default struct {
	access string
	secret string
}

func NewDefault(access string, secret string) *Default {
	return &Default{
		access: access,
		secret: secret,
	}
}

func (d *Default) Retrieve(_ context.Context) (credential aws.Credentials, err error) {
	credential = aws.Credentials{}
	credential.AccessKeyID = d.access
	credential.SecretAccessKey = d.secret
	credential.CanExpire = true
	credential.Expires = time.Date(9999, time.December, 31, 23, 59, 59, 999999999, time.UTC)

	return
}
