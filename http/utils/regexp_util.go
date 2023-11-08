package utils

import (
	"os"
	"regexp"
	"strings"
)

// 正则替换
func Replace(content *string) {
	re := regexp.MustCompile(`\${([^}]+)}`)
	matches := re.FindAllStringSubmatch(*content, -1)
	if matches != nil {
		for _, match := range matches {
			if len(match) >= 2 {
				value := os.Getenv(match[1])
				// 将value赋值给match[0]
				*content = strings.Replace(*content, match[0], value, -1)
			}
		}
	}
}
