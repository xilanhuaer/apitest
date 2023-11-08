package utils

import (
	"fmt"
	"reflect"
)

func InterfaceToString(origin interface{}) (data string, err error) {
	switch reflect.TypeOf(origin).Kind() {
	case reflect.String:
		return origin.(string), nil
	default:
		return fmt.Sprintf("%v", origin), nil
	}
}
