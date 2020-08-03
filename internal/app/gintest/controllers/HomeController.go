package controllers

import (
	"gintest/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewHomeController() *HomeController {
	return &HomeController{}
}

type HomeController struct {
	BaseController
}

func (controller HomeController) Ping(context *gin.Context) {
	logger.Info("ping~~~", logger.DEFAULT)
	context.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
