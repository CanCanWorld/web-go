package model

import "web-go/model/table"

func migration() {
	// 自动迁移
	DB.Set("gorm:table_option", "charset=utf8mb4").
		AutoMigrate(&table.User{}).
		AutoMigrate(&table.Module{}).
		AutoMigrate(&table.Article{})
}
