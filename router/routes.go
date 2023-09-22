package router

import (
	"demo/controllers"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.Use(Cors())
	r.POST("/upload", controllers.UploadControllersInstance.Upload) //文件上传
	r.POST("/user", controllers.UserControllersInstance.AddUser)    //新增
	r.GET("/user", controllers.UserControllersInstance.GetUser)     //查询
	r.PUT("/user", controllers.UserControllersInstance.Update)      //修改

	return r
}
