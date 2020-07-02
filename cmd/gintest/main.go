package main

import (
    "gintest/configs"
    "gintest/internal/app/gintest/routes"
    // "gintest/pkg/logger"
    "github.com/gin-gonic/gin"
)

func main() {

    // 設定環境變數
    configs.Setting()

    // gin 框架載入
    server := gin.Default()

    // 設定路由器
    routes.Setting(server)

    // 啟動 http server
    server.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
