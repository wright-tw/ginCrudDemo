package main

import (
	"gintest/configs"
	"gintest/internal/app/gintest/provider"
	"gintest/pkg/logger"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	configs.Setting()
	logger.Info("設定環境變數完成", logger.SERVER)

	server := gin.Default()
	logger.Info("gin 框架載入完成", logger.SERVER)

	ContainerProvider := provider.ContainerProvider{}
	router := ContainerProvider.GetInjectedRouter()
	router.Setting(server)
	logger.Info("路由注入設定完成", logger.SERVER)

	logger.Info("啟動 http server", logger.SERVER)
	runError := server.Run(":" + os.Getenv("APP_PORT"))
	if runError != nil {
		panic(runError)
	}
}
