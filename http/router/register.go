package router

import "github.com/gin-gonic/gin"

func Register(route *gin.Engine) {
	userGroup := route.Group("/user")
	{
		userGroup.POST("/register", userApi.Register)
		userGroup.POST("/login", userApi.Login)
		userGroup.GET("", userApi.List)
	}
}
