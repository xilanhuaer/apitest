package utils

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func StringToMap(data string) (header map[string]string, err error) {
	err = json.Unmarshal([]byte(data), &header)
	if err != nil {
		return nil, err
	}
	return header, nil
}

func InterfaceToString(origin interface{}) (data string, err error) {
	switch reflect.TypeOf(origin).Kind() {
	case reflect.String:
		return origin.(string), nil
	default:
		return fmt.Sprintf("%v", origin), nil
	}
}
