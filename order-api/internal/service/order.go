package service

import (
	"github.com/ANB98prog/order-api/internal/domain/entity"
	"time"
)

type CreateOrder struct {
	UserId   uint
	Products OrderItems
}

type OrderItem struct {
	ProductId uint
	Quantity  uint
}

type OrderItems []OrderItem

type Order struct {
	Id         uint
	ProductIds []uint
}

func (c *CreateOrder) ToDbOrder(orderDt time.Time) entity.Order {
	items := make([]entity.OrderItem, len(c.Products))
	for i, product := range c.Products {
		items[i] = entity.OrderItem{ProductID: product.ProductId, Quantity: product.Quantity}
	}

	return entity.Order{
		UserID: c.UserId,
		Date:   orderDt,
		Items:  items,
	}
}

func (items OrderItems) GetProductIds() []uint {
	ids := make([]uint, len(items))
	for i, item := range items {
		ids[i] = item.ProductId
	}
	return ids
}
