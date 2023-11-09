package main

import (
	"fmt"
	"os"
	"reflect"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/xilanhuaer/http-client/global"
	"github.com/xilanhuaer/http-client/middleware"
	"github.com/xilanhuaer/http-client/router"
	"github.com/xilanhuaer/http-client/utils"
)

func main() {
	// 生成公钥、私钥，只需要运行一次
	// utils.GenerateRSAKey(2048)
	// 读取配置文件，设置环境变量
	global.GetConfig("config.yaml")
	utils.SetEnv()
	// 打印设置的环境变量
	typeEnv := reflect.TypeOf(global.Config.Env)
	for i := 0; i < typeEnv.NumField(); i++ {
		field := typeEnv.Field(i)
		fmt.Println(field.Name, os.Getenv(field.Name))
	}
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(cors.Default())
	r.Use(middleware.JWTAuthMiddleware())
	router.Register(r)
	r.Run(":8080")
}
