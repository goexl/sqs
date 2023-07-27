package config

type Credentials struct {
	// 授权，相当于用户名
	Access string `json:"access" yaml:"access" xml:"access" toml:"access" validate:"required"`
	// 授权，相当于密码
	Secret string `json:"secret" yaml:"secret" xml:"secret" toml:"secret" validate:"required"`
}
