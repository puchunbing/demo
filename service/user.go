package services

import (
	"demo/config"
	"demo/model"
)

type InfoService struct{}

var (
	InfoServiceDao *InfoService
)

func init() {
	InfoServiceDao = &InfoService{}
}
func (c *InfoService) Add(v *model.UserInfo) error {
	o := config.GetMySql()
	info := model.UserInfo{}
	return info.Add(v, o)
}
