package main

import (
	"github.com/xilanhuaer/http-client/global"
	"gorm.io/gen"
)

func main() {
	global.GetConfig("../../config.yaml")
	global.GetConn()
	g := gen.NewGenerator(gen.Config{
		OutPath:      "../../dal/query",
		ModelPkgPath: "../../dal/model",
		Mode:         gen.WithQueryInterface | gen.WithDefaultQuery,
	})
	g.UseDB(global.DB)
	g.ApplyBasic(g.GenerateAllTable()...)
	g.Execute()
}
