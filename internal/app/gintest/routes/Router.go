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

	var UserController controllers.UserController
	var HomeController controllers.HomeController

	(&provider.Container{}).Injecting(

		// 控制器
		&inject.Object{Value: &UserController},
		&inject.Object{Value: &HomeController},
	)

	// crud
	userGroup := server.Group("users", middlewares.Check)
	userGroup.GET("", middlewares.Check2, UserController.List)
}
