package cache

import (
	"context"
	"github.com/ANB98prog/order-api/internal/config"
	"github.com/redis/go-redis/v9"
)

func NewRedisClient(cfg config.CacheConfig) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.Db,
	})

	result := client.Ping(context.Background())
	if result.Err() != nil {
		panic(result.Err())
	}

	return client
}
