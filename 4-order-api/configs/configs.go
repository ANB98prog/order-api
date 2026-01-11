package configs

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Db   DbConfig
	Auth AuthConfig
}

type DbConfig struct {
	Dsn string
}

type AuthConfig struct {
	Secret string
}

func ReadEnvironmentVariables() {
	viper.SetConfigName(".env")
	viper.SetConfigType("dotenv")
	viper.AddConfigPath("./4-order-api")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Fatal error config env file: %s \n", err)
	}
}

func LoadConfig() *Config {
	return &Config{
		Db:   DbConfig{Dsn: viper.GetString("DSN")},
		Auth: AuthConfig{Secret: viper.GetString("AUTH_SECRET")},
	}
}
