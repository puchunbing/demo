package main

import (
	"demo/auto"
	"demo/config"
	"demo/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(router.Cors())
	router.CollectRoute(r)
	conf := config.GetConfig()
	r.Run(":" + conf.Httpport)
}

func init() {
	config.LoadConfig()   //初始化配置文件
	config.InitMySql()    //初始化MySQL
	go auto.AutoMigrate() //gorm 数据库迁移
	router.NewLogger()    //初始化日志
}
