package repository

import (
	goerrors "errors"
	"fmt"
	"github.com/ANB98prog/purple-school-homeworks/4-order-api/internal/domain/entity"
	"github.com/ANB98prog/purple-school-homeworks/4-order-api/pkg/db"
	"github.com/ANB98prog/purple-school-homeworks/4-order-api/pkg/errors"
	"gorm.io/gorm"
)

type OrderRepository interface {
	Create(order *entity.Order) error
	GetById(id, userId uint) (*entity.Order, error)
	GetUserOrders(userId uint) ([]entity.Order, error)
}

type orderRepository struct {
	db *db.Db
}

func NewOrderRepository(db *db.Db) OrderRepository {
	return &orderRepository{db: db}
}

var _ OrderRepository = (*orderRepository)(nil)

// Create - создает заказ
func (repo *orderRepository) Create(order *entity.Order) error {
	result := repo.db.Create(&order)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// GetById - возвращает заказ по идентификатору
func (repo *orderRepository) GetById(id, userId uint) (*entity.Order, error) {
	order := &entity.Order{}
	result := repo.db.
		Preload("Items.Product").
		First(&order, "id = ? and user_id = ?", id, userId)

	if result.Error != nil {
		if goerrors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.NewItemNotFound(fmt.Sprintf("order with id %d not found", id))
		}
		return nil, result.Error
	}

	return order, nil
}

// GetUserOrders - возвращает все заказы пользователя
func (repo *orderRepository) GetUserOrders(userId uint) ([]entity.Order, error) {
	var orders []entity.Order
	result := repo.db.
		Preload("Items.Product").
		Find(&orders, "user_id = ?", userId)
	if result.Error != nil {
		return nil, result.Error
	}

	return orders, nil
}
