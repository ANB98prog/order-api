package configs

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Db DbConfig
}

type DbConfig struct {
	Dsn string
}

func LoadConfig() *Config {
	err := godotenv.Load("./4-order-api/.env.Development")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		Db: DbConfig{Dsn: os.Getenv("DSN")},
	}
}
