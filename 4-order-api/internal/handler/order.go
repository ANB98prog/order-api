package handler

type CreateOrderRequest struct {
	Items []CreateOrderItem `json:"items"`
}

type CreateOrderItem struct {
	ProductId uint `json:"product_id"`
	Quantity  uint `json:"quantity"`
}

type OrderResponse struct {
	Id    uint   `json:"id"`
	Items []uint `json:"items"`
}
