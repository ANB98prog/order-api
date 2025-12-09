package conf

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	EmailSender EmailSenderConfig
}

type EmailSenderConfig struct {
	From       string
	ApiAddress string
	SmtpAuth   SmtpConfig
}

type SmtpConfig struct {
	Login    string
	Password string
	Host     string
	Port     string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file, using default config. Error %s\n", err)
	}

	return &Config{
		EmailSender: EmailSenderConfig{
			From:       os.Getenv("from"),
			ApiAddress: os.Getenv("api_address"),
			SmtpAuth: SmtpConfig{
				Login:    os.Getenv("smtp_login"),
				Password: os.Getenv("smtp_password"),
				Host:     os.Getenv("smtp_host"),
				Port:     os.Getenv("smtp_port"),
			},
		},
	}
}
