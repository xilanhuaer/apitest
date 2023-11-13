package utils

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/oliveagle/jsonpath"
)

func DataExtract(content, path, name string) (err error) {
	var jsonData interface{}

	err = json.Unmarshal([]byte(content), &jsonData)
	if err != nil {
		return err
	}
	lookup, err := jsonpath.JsonPathLookup(jsonData, path)
	if err != nil {
		return err
	}
	value := fmt.Sprintf("%v", lookup)
	err = os.Setenv(name, value)
	return err
}
