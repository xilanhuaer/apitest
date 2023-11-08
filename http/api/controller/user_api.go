package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/xilanhuaer/http-client/model/entity"
)

type UserApi struct{}

func (u *UserApi) Create(context *gin.Context) {
	var (
		user entity.User
		err  error
	)
	if err = context.ShouldBindJSON(&user); err != nil {
	}
}
