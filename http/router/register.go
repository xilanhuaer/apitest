package router

import "github.com/gin-gonic/gin"

func Register(route *gin.Engine) {
	userGroup := route.Group("/user")
	{
		userGroup.GET("", userApi.Create)
	}
}
