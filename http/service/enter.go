package service

import "github.com/xilanhuaer/http-client/service/impl"

type ServiceGroup struct {
	ServiceGroup impl.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
