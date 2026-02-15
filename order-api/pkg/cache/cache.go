package cache

import (
	"context"
	"github.com/ANB98prog/order-api/configs"
	"github.com/redis/go-redis/v9"
)

func NewRedisClient(config *configs.CacheConfig) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: config.Password,
		DB:       0,
	})

	result := client.Ping(context.Background())
	if result.Err() != nil {
		panic(result.Err())
	}

	return client
}
