package auto

import (
	"demo/config"
	"demo/model"
)

// 数据库迁移
func AutoMigrate() {
	db := config.GetMySql()
	db.AutoMigrate(new(model.Info))
}
