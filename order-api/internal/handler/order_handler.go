package handler

import (
	goerrors "errors"
	"github.com/ANB98prog/order-api/configs"
	"github.com/ANB98prog/order-api/internal/service"
	"github.com/ANB98prog/order-api/pkg/errors"
	"github.com/ANB98prog/order-api/pkg/middlewares"
	"github.com/ANB98prog/order-api/pkg/request"
	"github.com/ANB98prog/order-api/pkg/response"
	"net/http"
	"strconv"
)

type OrderHandlerDeps struct {
	service.OrderService
}

type OrderHandler struct {
	orderService service.OrderService
}

func NewOrderHandler(router *http.ServeMux, deps OrderHandlerDeps, config *configs.Config) {
	handler := &OrderHandler{orderService: deps.OrderService}

	// Routing
	router.Handle("POST /order/create", middlewares.Authorization(handler.createOrder(), config))
	router.Handle("GET /order/{id}", middlewares.Authorization(handler.getOrderById(), config))
	router.Handle("GET /my-orders", middlewares.Authorization(handler.getUserOrders(), config))
}

func (handler *OrderHandler) createOrder() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		payload, err := request.HandleBody[CreateOrderRequest](&w, r)
		if err != nil {
			return
		}

		userId, ok := r.Context().Value(middlewares.CtxUserId).(uint)
		if !ok {
			response.Unauthorized(w, response.ErrorMessage{Message: "user id not found"})
			return
		}

		products := make(service.OrderItems, len(payload.Items))
		for i, item := range payload.Items {
			products[i] = service.OrderItem{ProductId: item.ProductId, Quantity: item.Quantity}
		}

		createOrder := service.CreateOrder{
			UserId:   userId,
			Products: products,
		}

		result, err := handler.orderService.CreateOrder(createOrder)
		if err != nil {
			if goerrors.Is(err, &errors.ItemNotFound{}) {
				response.NotFound(w, response.ErrorMessage{Message: err.Error()})
			} else {
				response.InternalServerError(w, response.ErrorMessage{Message: err.Error()})
			}

			return
		}

		createdOrder := OrderResponse{
			Id:    result.Id,
			Items: result.ProductIds,
		}

		response.Created(w, createdOrder)
	}
}

func (handler *OrderHandler) getOrderById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		orderId, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			response.BadRequest(w, response.ErrorMessage{Message: "order id is not set"})
			return
		}

		userId, ok := r.Context().Value(middlewares.CtxUserId).(uint)
		if !ok {
			response.Unauthorized(w, response.ErrorMessage{Message: "user id not found"})
			return
		}

		order, err := handler.orderService.GetOrderById(uint(orderId), userId)
		if err != nil {
			if goerrors.Is(err, &errors.ItemNotFound{}) {
				response.NotFound(w, response.ErrorMessage{Message: err.Error()})
			} else {
				response.InternalServerError(w, response.ErrorMessage{Message: err.Error()})
			}

			return
		}

		orderResponse := OrderResponse{Id: order.Id, Items: order.ProductIds}
		response.OKWithData(w, orderResponse)
	}
}

func (handler *OrderHandler) getUserOrders() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId, ok := r.Context().Value(middlewares.CtxUserId).(uint)
		if !ok {
			response.Unauthorized(w, response.ErrorMessage{Message: "user id not found"})
			return
		}

		orders, err := handler.orderService.GetUserOrders(userId)
		if err != nil {
			response.InternalServerError(w, response.ErrorMessage{Message: err.Error()})
			return
		}

		orderResponse := make([]OrderResponse, len(orders))
		for i, order := range orders {
			orderResponse[i] = OrderResponse{Id: order.Id, Items: order.ProductIds}
		}

		response.OKWithData(w, orderResponse)
	}
}
