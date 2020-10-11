package main

import (
	"gintest/configs"
	"gintest/internal/app/gintest/database/mysql"
	"gintest/internal/app/gintest/models"
)

func main() {
	configs.Setting()
	db := mysql.NewDB()
	db.Connect.AutoMigrate(
		// 把需要建立表的模型全部丟進來
		&models.User{},
	)
}
