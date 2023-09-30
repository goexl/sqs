package test_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/goexl/sqs"
)

var client *sqs.Client

func TestMain(m *testing.M) {
	builder := sqs.New()
	builder = builder.Credential().Default(os.Getenv("AWS_ID"), os.Getenv("AWS_KEY")).Build()
	client = builder.Queue("default", "test").Region("ap-northeast-2").Build()
	m.Run()
}

func TestSend(t *testing.T) {
	user := new(User)
	user.Name = "store"

	// 发送普通消息
	if out, se := client.Send().Data(user).Build().Do(context.Background()); nil != se {
		t.FailNow()
	} else {
		t.Logf("普通消息发送成功，消息标识：%s", *out.MessageId)
	}

	// 发送延迟消息
	if out, se := client.Send().Data(user).Delay(time.Hour).Build().Do(context.Background()); nil != se {
		t.FailNow()
	} else {
		t.Logf("延迟消息发送成功，消息标识：%s", *out.MessageId)
	}

	// 发送固定时间消息
	if out, se := client.Send().Data(user).Fix(time.Now().Add(time.Hour)).Build().Do(context.Background()); nil != se {
		t.FailNow()
	} else {
		t.Logf("固定时间消息发送成功，消息标识：%s", *out.MessageId)
	}
}

func TestHandle(t *testing.T) {
	if he := client.Handle().Build().Start(context.Background(), new(Handler)); nil != he {
		t.FailNow()
	}
}
