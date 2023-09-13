package model

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
)

type Info struct {
	ID        uint   `gorm:"primarykey"`
	Name      string `gorm:"size(50)"`
	Address   string `gorm:"column(address);size(50)"`
	Phone     string `gorm:"unique;not null"`
	Sex       int    `gorm:"sex"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime `gorm:"index"`
}

func (Info) TableName() string {
	return "info"
}

func (Info) Add(v *Info, o *gorm.DB) error {
	result := o.Create(&v)
	return result.Error
}
