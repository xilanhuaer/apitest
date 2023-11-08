package global

import (
	"log"
	"os"

	"github.com/xilanhuaer/http-client/model/common/config"
	"gopkg.in/yaml.v3"
)

var Config config.Config

func GetConfig(path string) {
	configBytes, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf(err.Error())
	}
	err = yaml.Unmarshal(configBytes, &Config)
	if err != nil {
		log.Fatalf(err.Error())
	}
}
