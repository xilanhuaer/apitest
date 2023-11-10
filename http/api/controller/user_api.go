package controller

import (
	"fmt"
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
		response.FailWithMessage(err.Error(), context)
		return
	}
	if register.Code != os.Getenv("RegisterCode") {
		response.FailWithMessage("邀请码错误", context)
		return
	}
	user = register.User
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
	params := map[string]string{
		"name":    context.DefaultQuery("name", ""),
		"email":   context.DefaultQuery("email", ""),
		"account": context.DefaultQuery("account", ""),
	}
	limit, offset, err := utils.PageUtil(
		context.DefaultQuery("page", "1"),
		context.DefaultQuery("page_size", "10"),
	)
	if err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	data, err := userService.List(params, limit, offset)
	if err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	response.OKWithData(data, context)
}

// Find userinfo
func (u *UserApi) Find(context *gin.Context) {
	id := context.Param("id")
	userinfo, err := userService.Find(id)
	if err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	response.OKWithData(userinfo, context)
}

// UpdatePassword 修改密码
func (u *UserApi) UpdatePassword(context *gin.Context) {
	id := context.Param("id")
	userId, ok := context.MustGet("userId").(uint)
	if ok {
		if id != fmt.Sprintf("%v", userId) {
			response.FailWithMessage("你无权修改此账号的密码", context)
			return
		}
	} else {
		response.FailWithMessage("获取身份信息失败，请重新登录", context)
		return
	}
	var password struct {
		Old   string
		New   string
		Renew string
	}
	if err := context.ShouldBindJSON(&password); err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	err := userService.UpdatePassword(password.Old, password.New, id)
	if err != nil {
		response.FailWithMessage(err.Error(), context)
		return
	}
	response.OK(context)
}

// Update 更新用户信息
func (u *UserApi) Update(context *gin.Context) {
	id := context.Param("id")
	userId, ok := context.MustGet("userId").(uint)
	if ok {
		if id != fmt.Sprintf("%v", userId) {
			response.FailWithMessage("你无权修改此账号的信息", context)
			return
		}
	} else {
		response.FailWithMessage("获取身份信息失败，请重新登录", context)
		return
	}
	var message struct {
		Name        string
		Avatar      string
		Email       string
		Phone       string
		Description string
	}
	err := context.ShouldBindJSON(&message)
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
