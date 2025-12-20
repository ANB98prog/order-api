package logging

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"strings"
)

type LogConfig struct {
	AppName  string       `mapstructure:"app_name,omitempty"`
	Level    string       `mapstructure:"log_level"`
	Format   string       `mapstructure:"log_format"`
	LogHooks []HookConfig `mapstructure:"log_hooks"`
}

type HookConfig struct {
	Type string `mapstructure:"type"`
	Path string `mapstructure:"path"`
}

func ReadLogConfig() (*LogConfig, error) {
	env := viper.GetString("ENVIRONMENT")
	if env == "" {
		env = "development"
	}

	appName := strings.ToLower(viper.GetString("APP_NAME"))
	configName := fmt.Sprintf("%s.%s", appName, env)
	viper.SetConfigName(configName)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./4-order-api")

	// Дефолтные значения
	viper.SetDefault("LOG_LEVEL", "info")
	viper.SetDefault("LOG_FORMAT", "json")

	err := viper.ReadInConfig()
	if err != nil {
		// Обработка ошибки: логируем и продолжаем с дефолтами.
		log.Printf("Failed to read config file: %v. Using defaults.\n", err)
		return nil, err
	}

	var logConfig LogConfig
	if err = viper.Unmarshal(&logConfig); err != nil {
		log.Printf("Failed to unmarshal config file: %v. Using defaults.\n", err)
		return nil, err
	}

	if appName != "" {
		logConfig.AppName = appName
	}

	return &logConfig, nil
}
