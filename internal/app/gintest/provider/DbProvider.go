package provider

import (
	"fmt"
	"gintest/pkg/logger"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
)

var dbConnection *gorm.DB

func GetDbConnection() *gorm.DB {

	if dbConnection != nil {
		return dbConnection
	}

	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	dbName := os.Getenv("MYSQL_DB")

	dbInfoString := fmt.Sprintf("%s:%s@(%s:%s)/%s", user, password, host, port, dbName)
	connection, dbConErr := gorm.Open("mysql", dbInfoString)

	if dbConErr != nil {
		logger.Error(dbConErr, logger.SERVER)
		panic(dbConErr)
	}

	if connection.Error != nil {
		logger.Error(connection.Error, logger.SERVER)
		panic(connection.Error)
	}

	dbConnection = connection

	return dbConnection
}
