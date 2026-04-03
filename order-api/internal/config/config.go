package config

import (
	"os"
	"strconv"
)

const defaultAddr = ":8081"

// Config конфигурация приложения
type Config struct {
	Addr  string
	Db    DbConfig
	Cache CacheConfig
	Auth  AuthConfig
}

// DbConfig конфигурация подключения к БД
type DbConfig struct {
	Dsn string
}

// CacheConfig конфигурация кеша
type CacheConfig struct {
	Addr     string
	Password string
	Db       int
}

// AuthConfig конфигурация авторизации
type AuthConfig struct {
	Secret string
}

// Load - Загружает конфиг приложения
func Load() Config {
	addr := getEnv("ADDR", defaultAddr)

	dbConfig := DbConfig{
		Dsn: getEnv("POSTGRES_DSN", ""),
	}

	cacheConfig := CacheConfig{
		Addr:     getEnv("REDIS_ADDR", "localhost:6379"),
		Password: getEnv("REDIS_PASS", ""),
		Db:       getEnvAsInt("REDIS_DB", 0),
	}

	authConfig := AuthConfig{
		Secret: getEnv("AUTH_SECRET", ""),
	}

	return Config{
		Addr:  addr,
		Db:    dbConfig,
		Cache: cacheConfig,
		Auth:  authConfig,
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok && value != "" {
		return value
	}

	return fallback
}

func getEnvAsInt(key string, fallback int) int {
	value, ok := os.LookupEnv(key)
	if !ok || value == "" {
		return fallback
	}

	parsed, err := strconv.Atoi(value)
	if err != nil {
		return fallback
	}

	return parsed
}
