package service

import (
	"github.com/ANB98prog/purple-school-homeworks/order-api/internal/repository"
)

type ProductService interface {
	Create(product CreateProduct) (Product, error)
	GetById(id uint) (Product, error)
	GetAll() ([]Product, error)
}

type productService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) ProductService {
	return &productService{repo: repo}
}

func (p *productService) Create(product CreateProduct) (Product, error) {
	dbProduct := product.ToDbProduct()
	err := p.repo.Create(dbProduct)
	if err != nil {
		return Product{}, err
	}

	createdProduct := FromDbProduct(dbProduct)

	return createdProduct, nil
}

func (p *productService) GetById(id uint) (Product, error) {
	dbProduct, err := p.repo.GetById(id)
	if err != nil {
		return Product{}, err
	}

	product := FromDbProduct(dbProduct)
	return product, nil
}

func (p *productService) GetAll() ([]Product, error) {
	dbProducts, err := p.repo.GetAll()
	if err != nil {
		return nil, err
	}

	products := make([]Product, len(dbProducts))
	for i, dbProduct := range dbProducts {
		products[i] = FromDbProduct(&dbProduct)
	}

	return products, nil
}
