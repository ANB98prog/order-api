package main

import (
	"github.com/ANB98prog/purple-school-homeworks/4-order-api/configs"
	"github.com/ANB98prog/purple-school-homeworks/4-order-api/pkg/logging"
)

func main() {
	configs.ReadEnvironmentVariables()
	//conf := configs.LoadConfig()
	config, err := logging.ReadLogConfig()
	if err != nil {
		panic(err)
	}
	logger := logging.ConfigureLogrus(config)

	//_ = db.NewDb(&conf.Db)

	logger.Info("First logging message")
}
