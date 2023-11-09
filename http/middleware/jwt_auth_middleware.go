package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/xilanhuaer/http-client/model/common/response"
	"github.com/xilanhuaer/http-client/utils"
)

func JWTAuthMiddleware() func(context *gin.Context) {
	return func(context *gin.Context) {
		path := context.Request.URL.Path
		switch path {
		case "/user/login", "/user/register":
			// context.Next()
			return
		}
		authHeader := context.Request.Header.Get("Authorization")
		if authHeader == "" {
			response.FailWithMessage("Authorization is empty, please login", context)
			return
		}
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			response.FailWithMessage("Not a standard token", context)
			return
		}
		mc, err := utils.ParseJWT(parts[1])
		if err != nil {
			response.FailWithMessage(err.Error(), context)
			return
		}
		context.Set("userId", mc.UserId)
		context.Set("userName", mc.UserName)
		context.Next()
	}
}
