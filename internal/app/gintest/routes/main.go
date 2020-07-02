package routes

import (
	"gintest/internal/app/gintest/controllers"
	"gintest/internal/app/gintest/middlewares"
	"github.com/gin-gonic/gin"
)

func Setting(server *gin.Engine) {

	HomeController := new(controllers.HomeController)

	server.GET("/ping", middlewares.Check, HomeController.Ping)

	// CRUD
	userGroup := server.Group("users", middlewares.Check)
	userGroup.PUT("/:id", middlewares.Check2, HomeController.Ping)
	userGroup.GET("", middlewares.Check2, HomeController.Ping)
	userGroup.POST("", middlewares.Check2, HomeController.Ping)
	userGroup.DELETE("/:id", middlewares.Check2, HomeController.Ping)
}
