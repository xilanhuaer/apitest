package config

type Env struct {
	Host         string `yaml:"host"`
	Issuer       string `yaml:"issuer"`
	RegisterCode string `yaml:"register_code"`
	UserSecret   string `yaml:"user_secret"`
}
