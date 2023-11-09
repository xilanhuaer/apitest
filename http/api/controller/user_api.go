package controller

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/xilanhuaer/http-client/model/common/response"
	"github.com/xilanhuaer/http-client/model/entity"
	"github.com/xilanhuaer/http-client/utils"
)

type UserApi struct{}

func (u *UserApi) Register(context *gin.Context) {
	var (
		user     entity.User
		err      error
		register struct {
			entity.User
			Code string `json:"code"`
		}
	)
	if err = context.ShouldBindJSON(&register); err != nil {
		response.OKWithMessage(err.Error(), context)
		return
	}
	if register.Code != os.Getenv("RegisterCode") {
		response.FailWithMessage("邀请码错误", context)
	}
	user = register.User
	if err = userService.Register(user); err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	response.OK(context)
}

func (u *UserApi) Login(context *gin.Context) {
	var (
		login struct {
			Account  string
			Password string
		}
		userinfo entity.UserInfo
	)
	if err := context.ShouldBindJSON(&login); err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	user, err := userService.Login(login.Account, login.Password)
	if err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	{
		userinfo.ID = user.ID
		userinfo.Name = user.Name
		userinfo.Account = user.Account
		userinfo.Email = user.Email
		userinfo.Phone = user.Phone
		userinfo.Description = user.Description
		userinfo.Token, _ = utils.GenJWT(user.ID, user.Name)
	}
	response.OKWithData(userinfo, context)
}

func (u *UserApi) List(context *gin.Context) {
}
