package provider

import (
	"gintest/internal/app/gintest/repositories"
	"gintest/internal/app/gintest/services"
	"gintest/pkg/logger"
	"github.com/facebookgo/inject"
)

type ContainerProvider struct{}

func (this *ContainerProvider) Injecting(objects ...*inject.Object) {

	var container inject.Graph

	// service
	objects = append(objects, &inject.Object{Value: &services.UserService{}})

	// repo
	objects = append(objects, &inject.Object{Value: &repositories.UserRepo{}})

	// DB 連線
	DbProvider := DbProvider{}
	DbConnection := DbProvider.GetDbConnection()
	objects = append(objects, &inject.Object{Value: DbConnection})

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
