package utils

import (
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/tidwall/gjson"
	"os"
)

func DataExtract(content, path, name string) {
	if err := os.Setenv(name, gconv.String(gjson.Get(content, path))); err != nil {
		panic(err)
	}
}
