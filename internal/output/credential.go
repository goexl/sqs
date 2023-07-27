package output

import (
	"time"
)

type Credential struct {
	// 授权，相当于用户名
	Id string `json:"id,omitempty"`
	// 授权，相当于密码
	Key string `json:"key,omitempty"`
	// 授权
	Session string `json:"session,omitempty"`
	// 过期时间
	Expires time.Time `json:"expires,omitempty"`
}
