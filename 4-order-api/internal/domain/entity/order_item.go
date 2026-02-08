package entity

type OrderItem struct {
	ID        uint    `gorm:"primary_key"`
	OrderID   uint    `gorm:"index"`
	ProductID uint    `gorm:"index"`
	Order     Order   `gorm:"foreignkey:OrderID"`
	Product   Product `gorm:"foreignkey:ProductID"`
	Quantity  uint
	Price     float64
}

type OrderItems []OrderItem

func (items OrderItems) GetProductIds() []uint {
	ids := make([]uint, len(items))
	for i, item := range items {
		ids[i] = item.Product.ID
	}

	return ids
}
