module gintest

// go 版本
go 1.14

// gin 框架
require github.com/gin-gonic/gin v1.6.3

// log 套件
require github.com/sirupsen/logrus v1.6.0

// gorm
require github.com/jinzhu/gorm v1.9.14

// 環境套件 .env
require github.com/joho/godotenv v1.3.0

require (
	// DI套件
	github.com/facebookgo/inject v0.0.0-20180706035515-f23751cae28b
	github.com/facebookgo/structtag v0.0.0-20150214074306-217e25fb9691 // indirect
)
