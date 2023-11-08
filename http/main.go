package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/xilanhuaer/http-client/global"
	"github.com/xilanhuaer/http-client/router"
	"github.com/xilanhuaer/http-client/utils"
)

func main() {
	global.GetConfig("config.yaml")
	utils.SetEnv()
	global.GetConn()

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(cors.Default())
	router.Register(r)
	r.Run(":8080")
}
