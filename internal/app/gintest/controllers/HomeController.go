package controllers

import (
	"gintest/pkg/logger"
	"github.com/gin-gonic/gin"
)

type HomeController struct{}

func (controller *HomeController) Ping(context *gin.Context) {
	logger.Info("ping~~~", "default")
	context.JSON(200, gin.H{
		"message": "pong",
	})
}
