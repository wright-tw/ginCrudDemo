package middlewares

import "github.com/gin-gonic/gin"

func Check2(context *gin.Context) {

	context.JSON(200, gin.H{
		"message": "請求處理前2",
	})

	context.Next()

	context.JSON(200, gin.H{
		"message": "請求處理後2",
	})
}
