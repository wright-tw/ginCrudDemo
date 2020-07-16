package routes

import (
	"gintest/internal/app/gintest/controllers"
	"gintest/internal/app/gintest/middlewares"
	"gintest/internal/app/gintest/provider"
	"github.com/facebookgo/inject"
	"github.com/gin-gonic/gin"
)

type Router struct{}

func (this *Router) Setting(server *gin.Engine) {

	var ContainerProvider provider.ContainerProvider
	var UserController controllers.UserController
	var HomeController controllers.HomeController

	ContainerProvider.Injecting(

		// 控制器
		&inject.Object{Value: &UserController},
		&inject.Object{Value: &HomeController},
	)

	// ping
	server.GET("ping", HomeController.Ping)

	// crud
	userGroup := server.Group("users", middlewares.Check)
	userGroup.GET("", middlewares.Check2, UserController.List)
	userGroup.PUT("/:id", middlewares.Check2, UserController.Update)
	userGroup.POST("", middlewares.Check2, UserController.Create)
	userGroup.DELETE("/:id", middlewares.Check2, UserController.Delete)
}
