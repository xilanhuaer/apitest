package utils

import (
	"encoding/json"
	"os"

	"github.com/oliveagle/jsonpath"
)

func DataExtract(content, path, name string) (err error) {
	var jsondata interface{}

	err = json.Unmarshal([]byte(content), &jsondata)
	if err != nil {
		return err
	}
	lookup, err := jsonpath.JsonPathLookup(jsondata, path)
	if err != nil {
		return err
	}
	value, err := InterfaceToString(lookup)
	if err != nil {
		return err
	}
	os.Setenv(name, value)
	return nil
}
