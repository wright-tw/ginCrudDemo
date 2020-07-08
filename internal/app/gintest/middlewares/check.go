package middlewares

import (
	"gintest/pkg/logger"
	"github.com/gin-gonic/gin"
)

func Check(context *gin.Context) {
	logger.Info("請求處理前", "middleware")

	context.Next()

	logger.Info("請求處理後", "middleware")
}
