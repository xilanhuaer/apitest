package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func result(code int, data interface{}, message string, context *gin.Context) {
	context.JSON(code, Response{
		Code:    code,
		Data:    data,
		Message: message,
	})
}

func OK(context *gin.Context) {
	result(http.StatusOK, nil, "success", context)
}

func OKWithMessage(message string, context *gin.Context) {
	result(http.StatusOK, nil, message, context)
}

func OKWithData(data interface{}, context *gin.Context) {
	result(http.StatusOK, data, "success", context)
}

func Fail(context *gin.Context) {
	result(http.StatusInternalServerError, nil, "fail", context)
}

func FailWithMessage(message string, context *gin.Context) {
	result(http.StatusInternalServerError, nil, message, context)
}
