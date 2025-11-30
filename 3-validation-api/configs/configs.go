package configs

import (
	"github.com/joho/godotenv"
	"log"
)

type Config struct {
	Email    EmailConfig
	Password PasswordConfig
	Address  AddressConfig
}

type EmailConfig struct {
	Auth              EmailAuthConfig
	EmailVerification EmailVerificationConfig
}

type EmailAuthConfig struct {
	Username string
	Password string
	Host     string
}

type EmailVerificationConfig struct {
	From    string
	Subject string
	Text    string
}

type PasswordConfig struct {
}

type AddressConfig struct {
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file, using default config")
	}

	return &Config{
		Email:    EmailConfig{},
		Password: PasswordConfig{},
		Address:  AddressConfig{},
	}
}
