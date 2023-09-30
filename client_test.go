package sqs_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/goexl/sqs"
)

type user struct {
	Name string `json:"name,omitempty"`
}

func TestSend(t *testing.T) {
	_user := new(user)
	_user.Name = "store"

	builder := sqs.New()
	builder = builder.Credential().Default(os.Getenv("AWS_ID"), os.Getenv("AWS_KEY")).Build()
	client := builder.Queue("default", "test").Region("ap-northeast-2").Build()

	// 发送普通消息
	if out, se := client.Send().Data(_user).Build().Do(context.Background()); nil != se {
		t.FailNow()
	} else {
		t.Logf("普通消息发送成功，消息标识：%s", *out.MessageId)
	}

	// 发送延迟消息
	if out, se := client.Send().Data(_user).Delay(time.Hour).Build().Do(context.Background()); nil != se {
		t.FailNow()
	} else {
		t.Logf("延迟消息发送成功，消息标识：%s", *out.MessageId)
	}

	// 发送固定时间消息
	if out, se := client.Send().Data(_user).Fix(time.Now().Add(time.Hour)).Build().Do(context.Background()); nil != se {
		t.FailNow()
	} else {
		t.Logf("固定时间消息发送成功，消息标识：%s", *out.MessageId)
	}
}
