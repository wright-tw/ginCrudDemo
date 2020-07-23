package routes

import (
	"gintest/internal/app/gintest/controllers"
	"gintest/internal/app/gintest/middlewares"
	"github.com/gin-gonic/gin"
)

func NewRouter(
	userController controllers.UserController,
	homeController controllers.HomeController,
) *Router {
	return &Router{
		UserController: userController,
		HomeController: homeController,
	}
}

type Router struct {
	UserController controllers.UserController
	HomeController controllers.HomeController
}

func (this *Router) Setting(server *gin.Engine) {

	// ping
	server.GET("ping", this.HomeController.Ping)

	// crud
	userGroup := server.Group("users", middlewares.Check)
	userGroup.GET("", middlewares.Check2, this.UserController.List)
	userGroup.PUT("/:id", middlewares.Check2, this.UserController.Update)
	userGroup.POST("", middlewares.Check2, this.UserController.Create)
	userGroup.DELETE("/:id", middlewares.Check2, this.UserController.Delete)
}
