package middlewares

import (
	"gintest/pkg/logger"
	"github.com/gin-gonic/gin"
)

func Check2(c *gin.Context) {
	logger.Info("請求處理前2", "middleware")
	c.Next()
	logger.Info("請求處理後2", "middleware")
}
