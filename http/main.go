package main

import (
	"github.com/xilanhuaer/http-client/global"
	"github.com/xilanhuaer/http-client/utils"
)

func main() {
	global.GetConfig("config.yaml")
	utils.SetEnv()
}
