package utils

import (
	"os"
	"reflect"

	"github.com/xilanhuaer/http-client/global"
)

func SetEnv() error {
	envType := reflect.TypeOf(global.Config.Env)
	for i := 0; i < envType.NumField(); i++ {
		field := envType.Field(i)
		err := os.Setenv(field.Name, reflect.ValueOf(global.Config.Env).Field(i).String())
		if err != nil {
			return err
		}
	}
	return nil
}
