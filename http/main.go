package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/xilanhuaer/http-client/global"
	"github.com/xilanhuaer/http-client/middleware"
	"github.com/xilanhuaer/http-client/router"
	"github.com/xilanhuaer/http-client/utils"
	"log"
	"os"
	"reflect"
)

func init() {
	global.GetConfig("config.yaml")
	global.GetConn()
	err := utils.SetEnv()
	if err != nil {
		log.Panic(err)
	}
	typeEnv := reflect.TypeOf(global.Config.Env)
	for i := 0; i < typeEnv.NumField(); i++ {
		field := typeEnv.Field(i)
		fmt.Println(field.Name, os.Getenv(field.Name))
	}
}

func main() {
	// 生成公钥、私钥，只需要运行一次
	// utils.GenerateRSAKey(2048)
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(cors.Default())
	r.Use(middleware.JWTAuthMiddleware())
	router.Register(r)
	_ = r.Run(":8080")
}
