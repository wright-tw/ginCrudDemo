package controllers

import (
	"gintest/pkg/logger"
	"github.com/gin-gonic/gin"
)

type HomeController struct {
	BaseController
}

func (controller *HomeController) Ping(context *gin.Context) {
	logger.Info("ping~~~", logger.DEFAULT)
	context.JSON(200, gin.H{
		"message": "pong",
	})
}
