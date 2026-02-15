package configs

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Db    DbConfig
	Cache CacheConfig
	Auth  AuthConfig
}

type DbConfig struct {
	Dsn string
}

type CacheConfig struct {
	Addr     string
	Password string
}

type AuthConfig struct {
	Secret string
}

func ReadEnvironmentVariables() {
	viper.SetConfigName(".env")
	viper.SetConfigType("dotenv")
	viper.AddConfigPath("./order-api")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Fatal error config env file: %s \n", err)
	}
}

func LoadConfig() *Config {
	return &Config{
		Db:    DbConfig{Dsn: viper.GetString("DSN")},
		Cache: CacheConfig{Addr: viper.GetString("REDIS_ADDR"), Password: viper.GetString("REDIS_PASSWORD")},
		Auth:  AuthConfig{Secret: viper.GetString("AUTH_SECRET")},
	}
}
