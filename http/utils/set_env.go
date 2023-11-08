package utils

import (
	"os"
	"reflect"

	"github.com/xilanhuaer/http-client/global"
)

func SetEnv() {
	envType := reflect.TypeOf(global.Config.Env)
	for i := 0; i < envType.NumField(); i++ {
		field := envType.Field(i)
		os.Setenv(field.Name, reflect.ValueOf(global.Config.Env).Field(i).String())
	}
}
