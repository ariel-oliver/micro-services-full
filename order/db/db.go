package db

import (
	"os"

	"github.com/redis/go-redis/v9"
)

var redisHost = os.Getenv("REDIS_HOST")

func Connect() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     redisHost + ":6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return client
}
