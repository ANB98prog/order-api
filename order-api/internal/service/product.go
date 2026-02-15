package service

import "github.com/ANB98prog/order-api/internal/domain/entity"

type CreateProduct struct {
	Name        string
	Description string
	Price       float64
}

type Product struct {
	Id          uint
	Name        string
	Description string
	Price       float64
}

func (p *CreateProduct) ToDbProduct() *entity.Product {
	return &entity.Product{
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
	}
}

func FromDbProduct(dbProduct *entity.Product) Product {
	return Product{
		Id:          dbProduct.ID,
		Name:        dbProduct.Name,
		Description: dbProduct.Description,
		Price:       dbProduct.Price,
	}
}
