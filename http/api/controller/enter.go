package controller

import "github.com/xilanhuaer/http-client/service"

type ApiGroup struct {
	UserApi
	SystemApi
}

var (
	userService   = service.ServiceGroupApp.ServiceGroup.UserService
	systemService = service.ServiceGroupApp.ServiceGroup.SystemService
)
