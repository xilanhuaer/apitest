package controller

import "github.com/xilanhuaer/http-client/service"

type ApiGroup struct {
	UserApi
}

var userService = service.ServiceGroupApp.ServiceGroup.UserService
