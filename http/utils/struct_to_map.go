package utils

import (
	"reflect"
	"strings"
)

// StructToMap 将结构体转为map，并设置key为小写
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
