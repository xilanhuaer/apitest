package router

import "github.com/gin-gonic/gin"

func Register(route *gin.Engine) {
	userGroup := route.Group("/user")
	{
		userGroup.POST("/register", userApi.Register)
		userGroup.POST("/login", userApi.Login)
		userGroup.GET("", userApi.List)
		userGroup.GET("/:id", userApi.Find)
		userGroup.PUT("/password/:id", userApi.UpdatePassword)
		userGroup.PUT(":/id", userApi.Update)
	}
}
