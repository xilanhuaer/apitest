package utils

import (
	"encoding/json"
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
	value, err := InterfaceToString(lookup)
	if err != nil {
		return err
	}
	err = os.Setenv(name, value)
	return err
}
