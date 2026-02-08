package handler

import (
	"fmt"
	"github.com/ANB98prog/purple-school-homeworks/4-order-api/configs"
	"github.com/ANB98prog/purple-school-homeworks/4-order-api/pkg/middlewares"
	"net/http"
)

type ProductHandler struct {
	*configs.Config
}

func NewProductHandler(router *http.ServeMux, config *configs.Config) {
	handler := &ProductHandler{config}

	router.Handle("POST /product/buy", middlewares.Authorization(handler.buyProduct(), config))
}

func (receiver ProductHandler) buyProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userPhone := r.Context().Value(middlewares.CtxUserPhone)

		fmt.Printf("User with phone %s want to buy product\n", userPhone)

		w.WriteHeader(http.StatusOK)
	}
}
