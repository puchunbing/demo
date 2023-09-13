package router

import (
	"demo/controllers"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.Use(Cors())
	r.POST("/upload", controllers.UploadControllersInstance.Upload) //文件上传

	return r
}
