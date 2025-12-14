package product

import (
	"errors"
	productErrors "github.com/ANB98prog/purple-school-homeworks/4-order-api/internal/product/product_errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ProductRepository interface {
	Create(product *Product) (*Product, error)
	Update(product *Product) (*Product, error)
	Delete(id uint) error
	GetById(id uint) (*Product, error)
	GetAll() ([]*Product, error)
}

type PgProductRepository struct {
	*gorm.DB
}

func (repo *PgProductRepository) Create(product *Product) (*Product, error) {
	result := repo.DB.Create(product)
	if err := result.Error; err != nil {
		return nil, err
	}

	return product, nil
}

func (repo *PgProductRepository) Update(product *Product) (*Product, error) {
	result := repo.DB.Clauses(clause.Returning{}).Updates(product)
	if result.Error != nil {
		return nil, result.Error
	}

	return product, nil
}

func (repo *PgProductRepository) Delete(id uint) error {
	result := repo.DB.Delete(&Product{}, id)
	return result.Error
}

func (repo *PgProductRepository) GetById(id uint) (*Product, error) {
	var product Product
	result := repo.DB.First(product, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, productErrors.ErrProductNotFound
		}
		return nil, result.Error
	}

	return &product, nil
}

func (repo *PgProductRepository) GetAll() ([]*Product, error) {
	var products []*Product
	result := repo.DB.Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}

	return products, nil
}
