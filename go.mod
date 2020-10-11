module gintest

// go version
go 1.14

// gin framework
require github.com/gin-gonic/gin v1.6.3

// log plugin
require github.com/sirupsen/logrus v1.6.0

// gorm
require github.com/jinzhu/gorm v1.9.14

// env plugin
require github.com/joho/godotenv v1.3.0

require (
	// redis plugin
	github.com/go-redis/redis/v8 v8.2.3
	// testing plugin
	github.com/stretchr/testify v1.6.1
	// DI plugin
	go.uber.org/dig v1.10.0
)
