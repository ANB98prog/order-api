package main

import (
	"github.com/ANB98prog/purple-school-homeworks/4-order-api/configs"
	"github.com/ANB98prog/purple-school-homeworks/4-order-api/internal/product"
	"github.com/ANB98prog/purple-school-homeworks/4-order-api/pkg/db"
)

func main() {
	conf := configs.LoadConfig()
	connection := db.NewDb(&conf.Db)
	err := connection.AutoMigrate(&product.Product{})
	if err != nil {
		panic(err)
	}
}
