package middlewares

import "github.com/gin-gonic/gin"

func Check(context *gin.Context) {

	context.JSON(200, gin.H{
		"message": "請求處理前",
	})

	context.Next()

	context.JSON(200, gin.H{
		"message": "請求處理後",
	})
}
