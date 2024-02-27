package main

import (
	"fmt"
	"os"
	"reflect"
	"sync"
	"time"

	"github.com/xilanhuaer/http-client/dal/query"
	"github.com/xilanhuaer/http-client/global"
	"github.com/xilanhuaer/http-client/utils"
)

func init() {
	global.GetConfig("config.yaml")
	global.GetConn()
	query.SetDefault(global.DB)
	utils.SetEnv()
	typeEnv := reflect.TypeOf(global.Config.Env)
	for i := 0; i < typeEnv.NumField(); i++ {
		field := typeEnv.Field(i)
		fmt.Println(field.Name, os.Getenv(field.Name))
	}
}

func main() {
	// 生成公钥、私钥，只需要运行一次
	// utils.GenerateRSAKey(2048)
	//r := gin.New()
	//r.Use(gin.Logger())
	//r.Use(cors.Default())
	//r.Use(middleware.JWTAuthMiddleware())
	//router.Register(r)
	//_ = r.Run(":8080")
	start := time.Now()
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_, err := utils.HTTP("GET", "/get", nil, nil, nil)
			if err != nil {
				return
			}
		}()
	}
	wg.Wait()
	fmt.Println(time.Since(start))
}
