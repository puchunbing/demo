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
		Pic:     "/imge/789.jpg",
		Address: "四川省成都市",
		Email:   "1982938@163.com",
		User: model.User{
			Age:  29,
			Name: "小涨",
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
	//var user model.User
	db := config.GetMySql()
	//db.Model(&userinfo).Association("users").Find(&userinfo.User)
	//查询info_id=1的数据放入userinfo中，并关联查询到User字段对应的数据
	db.Preload("User").Find(&userinfo, "info_id = ?", 1)
	//把user字段的数据放到user结构体中
	//db.Debug().Model(&userinfo).Related(&user,"User")
	msg := &logs.ResultInfo{}
	logs.Logger.Debug("挂了")
	logs.Logger.Info("第一级别")
	logs.Logger.Error("error")
	msg.Code = "0"
	msg.Message = "查询成功"
	msg.Data = userinfo
	c.JSONP(200, msg)
}

func (this *UserControllers) Update(c *gin.Context) {
	db := config.GetMySql()
	var userinfo model.UserInfo
	db.Preload("User").Find(&userinfo, "info_id =?", 1)
	db.Model(&userinfo.User).Update("age", 29)
	msg := &logs.ResultInfo{}
	msg.Code = "0"
	msg.Message = "修改成功"
	c.JSONP(200, msg)
}

func (this *UserControllers) Delete(c *gin.Context) {
	db := config.GetMySql()
	var userinfo model.UserInfo
	db.Preload("User").Find(&userinfo, "info_id =?", 1)
	//借助userinfo删除user信息，info数据不变
	//db.Delete(&userinfo.User)
	//删除userinfo
	db.Delete(&userinfo)
	msg := &logs.ResultInfo{}
	msg.Code = "0"
	msg.Message = "删除成功"
	c.JSONP(200, msg)
}
