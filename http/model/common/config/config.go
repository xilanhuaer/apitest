package config

type Config struct {
	Database Database `yaml:"database"`
	Env      Env      `yaml:"env"`
}
