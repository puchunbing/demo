package controllers

import (
	logs "demo/utils"
	"github.com/gin-gonic/gin"
)

type UploadControllers struct {
}

var (
	UploadControllersInstance *UploadControllers
)

func init() {
	UploadControllersInstance = &UploadControllers{}
}

func (this *UploadControllers) Upload(c *gin.Context) {
	f, err := c.FormFile("Filedata")
	if err != nil {
		msg := &logs.ResultInfo{}
		msg.Code = "-1"
		msg.Message = err.Error()
		c.JSONP(200, msg)
		return
	}
	err = c.SaveUploadedFile(f, "./imag/"+f.Filename)
	if err != nil {
		msg := &logs.ResultInfo{}
		msg.Code = "-1"
		msg.Message = err.Error()
		c.JSONP(200, msg)
		return
	}
	msg := &logs.ResultInfo{}
	msg.Code = "0"
	msg.Message = "上传成功"
	c.JSONP(200, msg)
}
