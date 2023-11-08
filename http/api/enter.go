package api

import "github.com/xilanhuaer/http-client/api/controller"

type ApiGroup struct {
	ApiGroup controller.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
