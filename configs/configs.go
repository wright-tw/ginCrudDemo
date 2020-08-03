package configs

import (
	_ "github.com/joho/godotenv/autoload" // init autoload
	"os"
)

func Setting() {
	// 設定時區
	os.Setenv("TZ", os.Getenv("APP_TIME_ZONE"))
}
