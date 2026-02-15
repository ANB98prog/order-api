package db

import (
	"github.com/ANB98prog/order-api/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Db struct {
	*gorm.DB
}

func NewDb(config *configs.DbConfig) *Db {
	dsn := config.Dsn
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		TranslateError: true,
	})
	if err != nil {
		panic(err)
	}

	return &Db{db}
}
