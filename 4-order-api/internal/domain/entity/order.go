package entity

import (
	"time"
)

type Order struct {
	ID     uint      `gorm:"primary_key"`
	UserID uint      `gorm:"index"`             // ForeignKey к User
	User   User      `gorm:"foreignKey:UserID"` // BelongsTo User
	Date   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	Total  float64
	Items  OrderItems `gorm:"foreignKey:OrderID"` // One-to-Many к OrderItem
}
