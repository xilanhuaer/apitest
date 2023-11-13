package utils

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func StringToInt32(data string) (int32, error) {
	id, err := strconv.Atoi(data)
	if err != nil {
		return 0, err
	}
	return int32(id), nil
}

func StringToMap(data string) (header map[string]string, err error) {
	err = json.Unmarshal([]byte(data), &header)
	if err != nil {
		return nil, err
	}
	return header, nil
}

func StructToMap(data interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	typ := reflect.TypeOf(data)
	val := reflect.ValueOf(data)
	for i := 0; i < typ.NumField(); i++ {
		fieldName := strings.ToLower(typ.Field(i).Name)
		fieldValue := val.Field(i).Interface()
		result[fieldName] = fieldValue
	}
	return result
}

func InterfaceToString(origin interface{}) (data string, err error) {
	switch reflect.TypeOf(origin).Kind() {
	case reflect.String:
		return origin.(string), nil
	default:
		return fmt.Sprintf("%v", origin), nil
	}
}
