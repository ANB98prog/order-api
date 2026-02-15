package entity

import (
	"github.com/lib/pq"
)

type Product struct {
	ID          uint   `gorm:"primary_key"`
	Name        string `gorm:"uniqueIndex"`
	Description string
	Price       float64
	Images      pq.StringArray `gorm:"type:text[]"`
	OrderItems  []OrderItem    `gorm:"many2many:order_items;"`
}
