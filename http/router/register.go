package router

import "github.com/gin-gonic/gin"

func Register(route *gin.Engine) {
	userGroup := route.Group("/user")
	{
		userGroup.POST("", userApi.Create)
		userGroup.GET("", userApi.List)
	}
}
