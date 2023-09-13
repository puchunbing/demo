package model

import "github.com/jinzhu/gorm"

type User struct {
	UserId int `gorm:"primary_key;AUTO_INCREMENT"`
	Age    int
	Name   string
	IID    int
}

type UserInfo struct {
	InfoId  int `gorm:"primary_key;AUTO_INCREMENT"`
	Pic     string
	Address string
	Email   string
	//关联关系
	User User `gorm:"ForeignKey:IID;AssociationForeignKey:InfoId"`
}

func (UserInfo) Add(v *UserInfo, o *gorm.DB) error {
	result := o.Create(&v)
	return result.Error
}
