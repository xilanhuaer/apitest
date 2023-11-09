package global

import (
	"fmt"
	"log"
	"os"

	"github.com/xilanhuaer/http-client/model/common/config"
	"gopkg.in/yaml.v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	Config      config.Config
	DB          *gorm.DB
	err         error
	configBytes []byte
)

// GetConfig 读取配置文件
func GetConfig(path string) {
	configBytes, err = os.ReadFile(path)
	if err != nil {
		log.Panic(err.Error())
	}
	err = yaml.Unmarshal(configBytes, &Config)
	if err != nil {
		log.Panic(err.Error())
	}
}

// GetConn 数据库连接
func GetConn() {
	database := Config.Database
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		database.User,
		database.Passwd,
		database.Host,
		database.Port,
		database.DB,
	)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		// 设置日志等级
		Logger: logger.Default.LogMode(logger.Info),
		// 缓存预编译语句
		PrepareStmt: true,
	})
	if err != nil {
		log.Panic(err.Error())
	}
}
