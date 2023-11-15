package main

import (
	"fmt"
	"log"
	"os"
	"reflect"

	"github.com/xilanhuaer/http-client/dal/query"
	"github.com/xilanhuaer/http-client/global"
	"github.com/xilanhuaer/http-client/utils"
)

func init() {
	global.GetConfig("config.yaml")
	global.GetConn()
	query.SetDefault(global.DB)
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
	//r := gin.New()
	//r.Use(gin.Logger())
	//r.Use(cors.Default())
	//r.Use(middleware.JWTAuthMiddleware())
	//router.Register(r)
	//_ = r.Run(":8080")
	//maps := map[string]interface{}{}
	//data := make([]entity.Data, 0)
	//utils.ReadExcel("data.xlsx", &data, maps)
	//for _, v := range data {
	//	fmt.Println(v)
	//}

}
