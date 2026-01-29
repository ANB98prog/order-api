package cache

import (
	"github.com/ANB98prog/purple-school-homeworks/4-order-api/configs"
	"github.com/redis/go-redis/v9"
)

func NewRedisClient(config *configs.CacheConfig) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: config.Password,
		DB:       0,
	})
}
