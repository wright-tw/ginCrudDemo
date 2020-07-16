package provider

import (
	"fmt"
	"gintest/internal/app/gintest/repositories"
	"gintest/internal/app/gintest/services"
	"gintest/pkg/logger"
	"github.com/facebookgo/inject"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
)

type Container struct{}

func (this *Container) Injecting(objects ...*inject.Object) {

	var container inject.Graph

	// service
	objects = append(objects, &inject.Object{Value: &services.UserService{}})

	// repo
	objects = append(objects, &inject.Object{Value: &repositories.UserRepo{}})

	// DB
	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	dbName := os.Getenv("MYSQL_DB")
	dbInfoString := fmt.Sprintf("%s:%s@(%s:%s)/%s", user, password, host, port, dbName)
	Db, dbConErr := gorm.Open("mysql", dbInfoString)
	if dbConErr != nil {
		logger.Error(dbConErr, logger.SERVER)
		panic(dbConErr)
	}
	if Db.Error != nil {
		logger.Error(Db.Error, logger.SERVER)
		panic(Db.Error)
	}
	objects = append(objects, &inject.Object{Value: Db})

	// 控制器 + 底層載入
	for _, object := range objects {
		providerErr := container.Provide(object)
		if providerErr != nil {
			logger.Error(providerErr, logger.SERVER)
			panic(providerErr)
		}
	}

	populateErr := container.Populate()
	if populateErr != nil {
		logger.Error(populateErr, logger.SERVER)
		panic(populateErr)
	}
}
