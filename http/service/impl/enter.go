package impl

import "github.com/xilanhuaer/http-client/utils"

type ServiceGroup struct {
	UserService
	SystemService
}

var exchange = utils.Exchange{}
