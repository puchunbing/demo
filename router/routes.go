package router

import (
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.Use(Cors())

	//var result Result
	//db := config.GetMySql()
	//db.Table("info").First(&result)
	//r.POST("/hello", controller.InfoControllerInstance.Add)

	return r
}
