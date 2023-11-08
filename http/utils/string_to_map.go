package utils

import "encoding/json"

func StringToMap(data string) (header map[string]string, err error) {
	err = json.Unmarshal([]byte(data), &header)
	if err != nil {
		return nil, err
	}
	return header, nil
}
