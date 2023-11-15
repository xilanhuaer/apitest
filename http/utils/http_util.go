package utils

import (
	"fmt"
	"github.com/gogf/gf/v2/util/gconv"
	"io"
	"net/http"
	"os"
	"strings"
)

// http请求相关方法

func DoHttp(method, path, params, headers, body string) (string, error) {
	var (
		request *http.Request
		err     error
		prefix  = os.Getenv("Prefix")
	)
	switch method {
	case "GET":
		request, err = http.NewRequest(method, prefix+path, nil)
	case "POST":
		request, err = http.NewRequest(method, prefix+path, strings.NewReader(body))
	case "PUT":
		request, err = http.NewRequest(method, prefix+path, strings.NewReader(body))
	default:
		request, err = http.NewRequest(method, prefix, nil)
	}
	if err != nil {
		return "", err
	}
	if params != "" {
		setParams(request, params)
	}
	if headers != "" {
		setHeaders(request, headers)
	}
	client := &http.Client{}
	res, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("HTTP request failed with status: %s", res.Status)
	}
	resByte, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(resByte), nil
}

// set header
func setHeaders(request *http.Request, headers string) error {
	Replace(&headers)
	header := gconv.Map(headers)
	for key, value := range header {
		request.Header.Set(key, gconv.String(value))
	}
	return nil
}

// set params
func setParams(request *http.Request, params string) {
	Replace(&params)
	request.URL.RawQuery = params
}
