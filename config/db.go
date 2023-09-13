package config

import (
	logs "demo/utils"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	msql *gorm.DB
)

func InitMySql() *gorm.DB {
	conf := GetConfig()
	dsn := conf.MySql.DbUser + `:` + conf.MySql.DbPass + `@tcp(` + conf.MySql.DbUrl + `)/` + conf.MySql.DbName + `?charset=utf8mb4&parseTime=True&loc=Local`
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		logs.Logger.Debug("数据库连接失败")
		return nil
	}
	msql = db
	return msql
}

func GetMySql() *gorm.DB {
	return msql
}
