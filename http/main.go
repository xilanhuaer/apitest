package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/xilanhuaer/http-client/router"
)

func main() {
	// 生成公钥、私钥，只需要运行一次
	// utils.GenerateRSAKey(2048)
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(cors.Default())
	router.Register(r)
	r.Run(":8080")
}
