package config

import (
	"time"
)

type Sqs struct {
	// 区域
	Region string `default:"ap-east-1" json:"region" yaml:"region" xml:"region" toml:"region" validate:"required"`
	// 授权
	Credentials Credentials `json:"credentials" yaml:"credentials" xml:"credentials" toml:"credentials" validate:"structonly"`
	// 队列
	Queue string `json:"queue" yaml:"queue" xml:"queue" toml:"queue" validate:"required_without=Queues"`
	// 队列列表
	Queues []Queue `json:"queues" yaml:"queues" xml:"queues" toml:"queues" validate:"required_without=Queue"`
	// 授权验证（策略）
	Session string `json:"session" yaml:"session" xml:"session" toml:"session" validate:"omitempty"`
	// 长轮询，减少费用
	Wait time.Duration `json:"wait" yaml:"wait" xml:"wait" toml:"wait" validate:"omitempty,max=20000"`
}
