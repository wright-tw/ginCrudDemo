package provider

import (
	"gintest/internal/app/gintest/controllers"
	"gintest/internal/app/gintest/models"
	"gintest/internal/app/gintest/repositories"
	"gintest/internal/app/gintest/routes"
	"gintest/internal/app/gintest/services"
	"go.uber.org/dig"
)

type ContainerProvider struct{}

var serviceList = []interface{}{

	// 控制器
	controllers.NewUserController,
	controllers.NewHomeController,

	// 服務
	services.NewUserService,

	// 資料組合
	repositories.NewUserRepo,

	// 模型
	models.NewUser,

	// 底層基本
	routes.NewRouter,
	GetDbConnection,
}

func (this *ContainerProvider) GetInjectedRouter() *routes.Router {

	container := dig.New()

	for _, function := range serviceList {
		if provideErr := container.Provide(function); provideErr != nil {
			panic(provideErr)
		}
	}

	// 自動注入所有控制器
	router := &routes.Router{}
	invokeErr := container.Invoke(func(readyRouter *routes.Router) {
		router = readyRouter
	})
	if invokeErr != nil {
		panic(invokeErr)
	}

	return router
}
