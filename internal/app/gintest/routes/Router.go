package routes

import (
	"gintest/internal/app/gintest/controllers"
	"gintest/internal/app/gintest/middlewares"
	"github.com/gin-gonic/gin"
)

func NewRouter(
	userController *controllers.UserController,
	homeController *controllers.HomeController,
) *Router {
	return &Router{
		UserController: userController,
		HomeController: homeController,
	}
}

type Router struct {
	UserController *controllers.UserController
	HomeController *controllers.HomeController
}

func (r *Router) Setting(server *gin.Engine) {
	// ping
	server.GET("ping", r.HomeController.Ping)

	// crud
	userGroup := server.Group("users", middlewares.Check)
	userGroup.GET("", middlewares.Check2, r.UserController.List)
	userGroup.PUT("/:id", middlewares.Check2, r.UserController.Update)
	userGroup.POST("", middlewares.Check2, r.UserController.Create)
	userGroup.DELETE("/:id", middlewares.Check2, r.UserController.Delete)
}
