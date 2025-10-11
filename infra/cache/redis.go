package cache

import (
	"context"
	"os"
	"sync"

	"github.com/redis/go-redis/v9"
)

var (
	RedisClient *redis.Client
	once        sync.Once
)

// InitRedis initializes the singleton Redis client instance
func InitRedis() error {
	var err error
	once.Do(func() {
		host := os.Getenv("REDIS_HOST")
		if host == "" {
			host = "localhost"
		}
		port := os.Getenv("REDIS_PORT")
		if port == "" {
			port = "6379"
		}
		password := os.Getenv("REDIS_PASSWORD")

		RedisClient = redis.NewClient(&redis.Options{
			Addr:     host + ":" + port,
			Password: password,
			DB:       0, // use default DB
		})

		// Test connection
		_, err = RedisClient.Ping(context.Background()).Result()
	})
	return err
}

// Close closes the Redis connection
func Close() error {
	if RedisClient != nil {
		return RedisClient.Close()
	}
	return nil
}
