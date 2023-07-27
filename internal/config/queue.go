package config

type Queue struct {
	// 名称
	Name string `json:"name" yaml:"name" xml:"name" toml:"name" validate:"required"`
	// 标签
	Label string `json:"label" yaml:"label" xml:"label" toml:"label" validate:"required"`
}
