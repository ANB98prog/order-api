package repository

import (
	goerr "errors"
	"fmt"
	"github.com/ANB98prog/purple-school-homeworks/4-order-api/internal/domain/entity"
	"github.com/ANB98prog/purple-school-homeworks/4-order-api/pkg/db"
	"github.com/ANB98prog/purple-school-homeworks/4-order-api/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ProductRepository interface {
	Create(product *entity.Product) error
	Update(product *entity.Product) error
	Delete(id uint) error
	GetById(id uint) (*entity.Product, error)
	Get(ids []uint) ([]entity.Product, error)
}

type pgProductRepository struct {
	*db.Db
}

func NewProductRepository(db *db.Db) ProductRepository {
	return &pgProductRepository{db}
}

var _ ProductRepository = (*pgProductRepository)(nil)

func (repo *pgProductRepository) Create(product *entity.Product) error {
	result := repo.DB.Create(product)
	if err := result.Error; err != nil {
		return err
	}

	return nil
}

func (repo *pgProductRepository) Update(product *entity.Product) error {
	result := repo.DB.Clauses(clause.Returning{}).Updates(product)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repo *pgProductRepository) Delete(id uint) error {
	result := repo.DB.Delete(&entity.Product{}, id)
	return result.Error
}

func (repo *pgProductRepository) GetById(id uint) (*entity.Product, error) {
	var product entity.Product
	result := repo.DB.Take(&product, id)
	if result.Error != nil {
		if goerr.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.NewItemNotFound(fmt.Sprintf("product with id %v not found", id))
		}
		return nil, result.Error
	}

	return &product, nil
}

func (repo *pgProductRepository) Get(ids []uint) ([]entity.Product, error) {
	var products []entity.Product

	query := repo.DB
	if len(ids) != 0 {
		query = query.Where("id IN ?", ids)
	}

	result := query.Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}

	return products, nil
}
