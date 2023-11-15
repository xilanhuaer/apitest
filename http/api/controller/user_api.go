package controller

import (
	"os"
	"strconv"

	"github.com/xilanhuaer/http-client/common/response"
	"github.com/xilanhuaer/http-client/dal/model"

	"github.com/gin-gonic/gin"
	"github.com/xilanhuaer/http-client/utils"
)

type UserApi struct{}

func (u *UserApi) Register(context *gin.Context) {
	var (
		user     model.User
		err      error
		register struct {
			Account     string `json:"account"`
			Password    string `json:"password"`
			Name        string `json:"name"`
			Avatar      string `json:"avatar"`
			Email       string `json:"email"`
			Phone       string `json:"phone"`
			Description string `json:"description"`
			Code        string `json:"code"`
		}
	)
	if err = context.ShouldBindJSON(&register); err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	if register.Code != os.Getenv("RegisterCode") {
		response.FailWithMessage("邀请码错误", context)
		return
	}
	// 拷贝register到user
	{
		user.Account = register.Account
		user.Name = register.Name
		user.Avatar = register.Avatar
		user.Password = register.Password
		user.Email = register.Email
		user.Phone = register.Phone
		user.Description = register.Description
	}
	if err = userService.Register(user); err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	response.OK(context)
}

func (u *UserApi) Login(context *gin.Context) {
	var login struct {
		Account  string
		Password string
	}
	if err := context.ShouldBindJSON(&login); err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	userinfo, err := userService.Login(login.Account, login.Password)
	if err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	response.OKWithData(userinfo, context)
}

// List 用户列表
func (u *UserApi) List(context *gin.Context) {
	params := context.Request.URL.RawQuery
	data, err := userService.List(params)
	if err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	response.OKWithData(data, context)
}

// Find userinfo
func (u *UserApi) Find(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	userinfo, err := userService.Find(id)
	if err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	response.OKWithData(userinfo, context)
}

// UpdatePassword 修改密码
func (u *UserApi) UpdatePassword(context *gin.Context) {
	id, err := utils.ValidUserAuthority(context)
	if err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	var password struct {
		Old   string
		New   string
		Renew string
	}
	if err = context.ShouldBindJSON(&password); err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	err = userService.UpdatePassword(password.Old, password.New, id)
	if err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	response.OK(context)
}

// Update 更新用户信息
func (u *UserApi) Update(context *gin.Context) {
	id, err := utils.ValidUserAuthority(context)
	if err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	var message struct {
		Name        string
		Avatar      string
		Email       string
		Phone       string
		Description string
	}
	err = context.ShouldBindJSON(&message)
	if err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	err = userService.Update(id, message)
	if err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
}
