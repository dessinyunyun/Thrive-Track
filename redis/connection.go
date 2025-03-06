package redis

import (
	"context"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

func NewRedisClient(log *logrus.Entry) *redis.Client {

	client := redis.NewClient(&redis.Options{
		Addr:         os.Getenv("REDIS_HOST"),
		MinIdleConns: 100,
		PoolSize:     100,
		PoolTimeout:  time.Duration(100) * time.Minute,
		Password:     "", // no password set
		DB:           0,  // use default DB
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Errorf("Failed to connect to Redis: %v", err)
	}

	return client
}
