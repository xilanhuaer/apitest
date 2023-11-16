package utils

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"os"
	"strings"
)

// import (
//
//	"fmt"
//	"github.com/gogf/gf/v2/util/gconv"
//	"io"
//	"net/http"
//	"os"
//	"strings"
//
// )
//
// // http请求相关方法
//
//	func DoHttp(method, path, params, headers, body string) (string, error) {
//		var (
//			request *http.Request
//			err     error
//			prefix  = os.Getenv("Prefix")
//		)
//		switch method {
//		case "GET":
//			request, err = http.NewRequest(method, prefix+path, nil)
//		case "POST":
//			request, err = http.NewRequest(method, prefix+path, strings.NewReader(body))
//		case "PUT":
//			request, err = http.NewRequest(method, prefix+path, strings.NewReader(body))
//		default:
//			request, err = http.NewRequest(method, prefix, nil)
//		}
//		if err != nil {
//			return "", err
//		}
//		if params != "" {
//			setParams(request, params)
//		}
//		if headers != "" {
//			setHeaders(request, headers)
//		}
//		client := &http.Client{}
//		res, err := client.Do(request)
//		if err != nil {
//			return "", err
//		}
//		defer res.Body.Close()
//		if res.StatusCode != http.StatusOK {
//			return "", fmt.Errorf("HTTP request failed with status: %s", res.Status)
//		}
//		resByte, err := io.ReadAll(res.Body)
//		if err != nil {
//			return "", err
//		}
//		return string(resByte), nil
//	}
//
// // set header
//
//	func setHeaders(request *http.Request, headers string) error {
//		Replace(&headers)
//		header := gconv.Map(headers)
//		for key, value := range header {
//			request.Header.Set(key, gconv.String(value))
//		}
//		return nil
//	}
//
// // set params
//
//	func setParams(request *http.Request, params string) {
//		Replace(&params)
//		request.URL.RawQuery = params
//	}
const (
	GET    = "GET"
	POST   = "POST"
	PUT    = "PUT"
	DELETE = "DELETE"
)

var client *resty.Client

func init() {
	client = resty.New()
}

// HTTP http请求
func HTTP(method, path string, params, headers map[string]string, body interface{}) (*resty.Response, error) {
	host := os.Getenv("Host")
	if host == "" {
		return nil, fmt.Errorf("host is empty")
	}
	method = strings.ToUpper(method)
	switch method {
	case GET, POST, PUT, DELETE:
	default:
		return nil, fmt.Errorf("method is not support")
	}
	resp, err := client.R().
		SetQueryParams(params).
		SetHeaders(headers).
		SetBody(body).
		Execute(method, host+path)
	return resp, err
}
