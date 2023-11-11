package config

type Database struct {
	User   string `yaml:"user"`
	Passwd string `yaml:"passwd"`
	Host   string `yaml:"host"`
	Port   string `yaml:"port"`
	DB     string `yaml:"db"`
}
