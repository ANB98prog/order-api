package service

import (
	"fmt"
	"github.com/ANB98prog/purple-school-homeworks/order-api/internal/repository"
	"github.com/ANB98prog/purple-school-homeworks/order-api/pkg/errors"
	"time"
)

type OrderService interface {
	CreateOrder(order CreateOrder) (Order, error)
	GetUserOrders(id uint) ([]Order, error)
	GetOrderById(id, userId uint) (Order, error)
}

type orderService struct {
	orderRepo   repository.OrderRepository
	productRepo repository.ProductRepository
	userRepo    repository.UserRepository
}

func NewOrderService(orderRepo repository.OrderRepository, productRepo repository.ProductRepository, userRepo repository.UserRepository) OrderService {
	return &orderService{orderRepo: orderRepo, productRepo: productRepo, userRepo: userRepo}
}

var _ OrderService = (*orderService)(nil)

func (service *orderService) CreateOrder(order CreateOrder) (Order, error) {

	// Проверяем существование продуктов
	err := service.checkProductsExists(order.Products)
	if err != nil {
		return Order{}, err
	}

	dbOrder := order.ToDbOrder(time.Now())
	err = service.orderRepo.Create(&dbOrder)
	if err != nil {
		return Order{}, err
	}

	createdOrder := Order{Id: dbOrder.ID, ProductIds: order.Products.GetProductIds()}
	return createdOrder, nil
}

func (service *orderService) checkProductsExists(orderItems OrderItems) error {
	productIds := orderItems.GetProductIds()
	products, err := service.productRepo.Get(productIds)
	if err != nil {
		return err
	}

	idSet := make(map[uint]struct{})
	for _, product := range products {
		idSet[product.ID] = struct{}{}
	}

	for _, id := range productIds {
		if _, exists := idSet[id]; !exists {
			return errors.NewItemNotFound(fmt.Sprintf("product with id %d does not exist", id))
		}
	}

	return nil
}

func (service *orderService) GetUserOrders(userId uint) ([]Order, error) {
	dbOrders, err := service.orderRepo.GetUserOrders(userId)
	if err != nil {
		return nil, err
	}

	orders := make([]Order, len(dbOrders))
	for i, order := range dbOrders {
		orders[i] = Order{order.ID, order.Items.GetProductIds()}
	}

	return orders, nil
}

func (service *orderService) GetOrderById(id, userId uint) (Order, error) {
	dbOrder, err := service.orderRepo.GetById(id, userId)
	if err != nil {
		return Order{}, err
	}

	return Order{dbOrder.ID, dbOrder.Items.GetProductIds()}, nil
}
