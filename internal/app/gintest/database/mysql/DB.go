package mysql

import (
	"fmt"
	"gintest/pkg/logger"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // INIT MYSQL
	"os"
)

func NewDB() *DB {
	db := DB{}
	db.init()
	return &db
}

type DB struct {
	Connect *gorm.DB
}

func (d *DB) init() {
	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	dbName := os.Getenv("MYSQL_DB")

	dbInfoString := fmt.Sprintf("%s:%s@(%s:%s)/%s?parseTime=true&loc=Local", user, password, host, port, dbName)
	connection, dbConErr := gorm.Open("mysql", dbInfoString)

	if dbConErr != nil {
		logger.Error(dbConErr, logger.DB)
	} else {
		logger.Info("資料庫連線成功1", logger.DB)
	}

	if connection.Error != nil {
		logger.Error(connection.Error, logger.DB)
	} else {
		logger.Info("資料庫連線成功2", logger.DB)
	}

	connection.LogMode(true)

	d.Connect = connection
}

func (d *DB) GetConnect() *gorm.DB {
	err := d.Connect.DB().Ping()
	if err == nil {
		return d.Connect
	}
	d.init()
	return d.Connect
}
