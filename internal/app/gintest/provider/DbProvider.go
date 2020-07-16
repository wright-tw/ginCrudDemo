package provider

import (
	"fmt"
	"gintest/pkg/logger"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
)

type DbProvider struct {
	dbConnection *gorm.DB
}

func (this *DbProvider) GetDbConnection() *gorm.DB {

	if this.dbConnection != nil {
		return this.dbConnection
	}

	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	dbName := os.Getenv("MYSQL_DB")

	dbInfoString := fmt.Sprintf("%s:%s@(%s:%s)/%s", user, password, host, port, dbName)
	dbConnection, dbConErr := gorm.Open("mysql", dbInfoString)

	if dbConErr != nil {
		logger.Error(dbConErr, logger.SERVER)
		panic(dbConErr)
	}

	if dbConnection.Error != nil {
		logger.Error(dbConnection.Error, logger.SERVER)
		panic(dbConnection.Error)
	}

	this.dbConnection = dbConnection

	return dbConnection
}
