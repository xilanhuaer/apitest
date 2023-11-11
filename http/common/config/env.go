package config

type Env struct {
	Prefix       string `yaml:"prefix"`
	Issuer       string `yaml:"issuer"`
	RegisterCode string `yaml:"register_code"`
	UserSecret   string `yaml:"user_secret"`
}
