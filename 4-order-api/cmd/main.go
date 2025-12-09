package main

import (
	"github.com/ANB98prog/purple-school-homeworks/4-order-api/configs"
	"github.com/ANB98prog/purple-school-homeworks/4-order-api/pkg/db"
)

func main() {
	conf := configs.LoadConfig()
	_ = db.NewDb(&conf.Db)
}
