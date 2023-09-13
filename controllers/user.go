package controllers

import (
	"demo/config"
	"demo/model"
	services "demo/service"
	logs "demo/utils"
	"github.com/gin-gonic/gin"
)

type UserControllers struct {
}

var (
	UserControllersInstance *UserControllers
)

func init() {
	UserControllersInstance = &UserControllers{}
}

func (this *UserControllers) AddUser(c *gin.Context) {

	user := &model.UserInfo{
		Pic:     "/imge/123.jpg",
		Address: "深圳市",
		Email:   "123@163.com",
		User: model.User{
			Age:  17,
			Name: "小李",
		},
	}
	err := services.InfoServiceDao.Add(user)
	if err != nil {
		msg := &logs.ResultInfo{}
		msg.Code = "-1"
		msg.Message = err.Error()
		c.JSONP(200, msg)
		return
	}
	msg := &logs.ResultInfo{}
	msg.Code = "0"
	msg.Message = "新增成功"
	c.JSONP(200, msg)
}

func (this *UserControllers) GetUser(c *gin.Context) {
	var userinfo model.UserInfo
	db := config.GetMySql()
	//db.Model(&userinfo).Association("users").Find(&userinfo.User)
	//查询info_id=1的数据放入userinfo中，并关联查询到User字段对应的数据
	db.Preload("User").Find(&userinfo, "info_id = ?", 1)
	msg := &logs.ResultInfo{}
	msg.Code = "0"
	msg.Message = "查询成功"
	msg.Data = userinfo
	c.JSONP(200, msg)
}
