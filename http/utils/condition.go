package utils

import (
	"fmt"
	"strings"
)

func ParseCondition(paramsString string) (map[string]interface{}, error) {
	// 如果没有条件，直接返回空字符串
	if paramsString == "" {
		return make(map[string]interface{}), nil
	}
	params := make(map[string]interface{})
	// 按照 # 分割条件
	conditions := strings.Split(paramsString, "#")
	// 遍历条件
	for _, condition := range conditions {
		// 按照 | 分割条件
		parts := strings.Split(condition, "|")
		// 如果条件格式不正确，返回错误
		if len(parts) == 3 {
			key := parts[0]
			values := parts[2]
			// 获取操作符，并转换为小写
			operation := strings.ToLower(parts[1])
			switch operation {
			case "eq":
				params[key] = values
			case "ne":
				params[key+" <>"] = values
			case "gt":
				params[key+" >"] = values
			case "ge":
				params[key+" >="] = values
			case "lt":
				params[key+" <"] = values
			case "le":
				params[key+" <="] = values
			case "like":
				params[key+" LIKE"] = values
			case "between":
				value := strings.Split(values, ",")
				if len(value) == 2 {
					params[key+" BETWEEN ? AND ?"] = value
				} else {
					return make(map[string]interface{}), fmt.Errorf("between operation needs two values")
				}
			default:
				return make(map[string]interface{}), fmt.Errorf("unsupported operation: %s", operation)
			}

		} else {
			return make(map[string]interface{}), fmt.Errorf("condition format error")
		}
	}

	return params, nil
}
