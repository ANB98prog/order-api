package main

import (
	"github.com/ANB98prog/purple-school-homeworks/4-order-api/configs"
	"github.com/ANB98prog/purple-school-homeworks/4-order-api/pkg/db"
	"github.com/ANB98prog/purple-school-homeworks/4-order-api/pkg/logging"
	"github.com/sirupsen/logrus"
)

func main() {
	configs.ReadEnvironmentVariables()
	conf := configs.LoadConfig()
	config, err := logging.ReadLogConfig()
	if err != nil {
		panic(err)
	}
	logging.ConfigureLogrus(config)

	_ = db.NewDb(&conf.Db)

	logrus.Debug("First logging message")
}
