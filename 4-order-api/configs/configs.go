package configs

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Db DbConfig
}

type DbConfig struct {
	Dsn string
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
		Db: DbConfig{Dsn: viper.GetString("DSN")},
	}
}
