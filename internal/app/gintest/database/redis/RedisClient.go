package redis

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"os"
	"strconv"
)

var redisClient *redis.Client

func GetRedisClient() *redis.Client {
	if redisClient != nil {
		return redisClient
	}

	hostAndPort := fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))
	dbName, _ := strconv.Atoi(os.Getenv("REDIS_DB"))
	client := redis.NewClient(&redis.Options{
		Addr:     hostAndPort,
		Password: os.Getenv("REDIS_PASSWORD"), // no password set
		DB:       dbName,                      // use default DB
	})

	redisClient = client

	return redisClient
}
