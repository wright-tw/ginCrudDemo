package provider

import (
	"gintest/internal/app/gintest/models"
	"gintest/internal/app/gintest/services"
	"gintest/pkg/logger"
	"github.com/facebookgo/inject"
)

type Container struct{}

func (this *Container) Injecting(objects ...*inject.Object) {

	var container inject.Graph

	// service
	objects = append(objects, &inject.Object{Value: &services.UserService{}})

	// model
	objects = append(objects, &inject.Object{Value: &models.User{}})

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
