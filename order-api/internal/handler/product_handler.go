package handler

import (
	goerrors "errors"
	"github.com/ANB98prog/purple-school-homeworks/order-api/configs"
	"github.com/ANB98prog/purple-school-homeworks/order-api/internal/service"
	"github.com/ANB98prog/purple-school-homeworks/order-api/pkg/errors"
	"github.com/ANB98prog/purple-school-homeworks/order-api/pkg/middlewares"
	"github.com/ANB98prog/purple-school-homeworks/order-api/pkg/request"
	"github.com/ANB98prog/purple-school-homeworks/order-api/pkg/response"
	"net/http"
	"strconv"
)

type ProductHandler struct {
	productService service.ProductService
}

func NewProductHandler(router *http.ServeMux, productService service.ProductService, config *configs.Config) {
	handler := &ProductHandler{
		productService: productService,
	}

	// Routing
	router.Handle("GET /products", middlewares.Authorization(handler.getProducts(), config))
	router.Handle("GET /product/{id}", middlewares.Authorization(handler.getProductById(), config))
	router.Handle("POST /product", middlewares.Authorization(handler.createProduct(), config))
}

func (h *ProductHandler) getProducts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		products, err := h.productService.GetAll()
		if err != nil {
			if goerrors.Is(err, &errors.ItemNotFound{}) {
				response.NotFound(w, response.ErrorMessage{Message: "product not found"})
				return
			}
			response.BadRequest(w, response.ErrorMessage{Message: err.Error()})
			return
		}

		result := make([]ProductResponse, len(products))
		for i, product := range products {
			result[i] = ProductResponse{
				Id:          product.Id,
				Name:        product.Name,
				Description: product.Description,
				Price:       product.Price,
			}
		}

		response.OKWithData(w, result)
	}
}

func (h *ProductHandler) getProductById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			response.BadRequest(w, response.ErrorMessage{Message: "product id is not set"})
			return
		}

		product, err := h.productService.GetById(uint(id))
		if err != nil {
			if goerrors.Is(err, &errors.ItemNotFound{}) {
				response.NotFound(w, response.ErrorMessage{Message: "product not found"})
				return
			}
			response.BadRequest(w, response.ErrorMessage{Message: err.Error()})
			return
		}

		result := ProductResponse{
			Id:          product.Id,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
		}

		response.OKWithData(w, result)
	}
}

func (h *ProductHandler) createProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		payload, err := request.HandleBody[CreateProductRequest](&w, r)
		if err != nil {
			return
		}

		createdProduct, err := h.productService.Create(service.CreateProduct{
			Name:        payload.Name,
			Description: payload.Description,
			Price:       payload.Price,
		})

		if err != nil {
			if goerrors.Is(err, errors.ErrItemAlreadyExists) {
				response.BadRequest(w, response.ErrorMessage{Message: err.Error()})
				return
			}

			response.InternalServerError(w, response.ErrorMessage{Message: err.Error()})
			return
		}

		result := ProductResponse{
			Id:          createdProduct.Id,
			Name:        createdProduct.Name,
			Description: createdProduct.Description,
			Price:       createdProduct.Price,
		}

		response.Created(w, result)
	}
}
