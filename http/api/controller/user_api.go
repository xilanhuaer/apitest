package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/xilanhuaer/http-client/model/common/response"
	"github.com/xilanhuaer/http-client/model/entity"
)

type UserApi struct{}

func (u *UserApi) Create(context *gin.Context) {
	var (
		user entity.User
		err  error
	)
	if err = context.ShouldBindJSON(&user); err != nil {
		response.OKWithMessage(err.Error(), context)
		return
	}
	if err = userService.Create(user); err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	response.OK(context)
}
