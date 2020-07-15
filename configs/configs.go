package configs

import (
	_ "github.com/joho/godotenv/autoload"
	"os"
)

func Setting() {

	// 設定時區
	os.Setenv("TZ", os.Getenv("TIME_ZONE"))

}