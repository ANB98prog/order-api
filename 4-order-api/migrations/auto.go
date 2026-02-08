package main

import (
	"github.com/ANB98prog/purple-school-homeworks/4-order-api/configs"
	"github.com/ANB98prog/purple-school-homeworks/4-order-api/internal/domain/entity"
	"github.com/ANB98prog/purple-school-homeworks/4-order-api/pkg/db"
)

func main() {
	configs.ReadEnvironmentVariables()
	conf := configs.LoadConfig()
	connection := db.NewDb(&conf.Db)

	if err := connection.AutoMigrate(&entity.User{}, &entity.Product{}, &entity.Order{}, &entity.OrderItem{}); err != nil {
		panic(err)
	}

	laptop := entity.Product{Name: "Laptop", Price: 999.99}
	mouse := entity.Product{Name: "Mouse", Price: 29.99}
	connection.Create(&laptop)
	connection.Create(&mouse)

	order := entity.Order{
		UserID: 1,
		Total:  1029.97,
		Items: []entity.OrderItem{
			{ProductID: laptop.ID, Quantity: 1, Price: 999.99},
			{ProductID: mouse.ID, Quantity: 1, Price: 29.99},
		},
	}
	result := connection.Create(&order)
	if result.Error != nil {
		panic(result.Error)
	}
}
