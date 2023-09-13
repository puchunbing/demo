package main

import (
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
	config.LoadConfig()
	config.InitMySql()
	//go auto.AutoMigrate()
}
