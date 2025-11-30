package verify

import (
	"github.com/ANB98prog/purple-school-homeworks/3-validation-api/configs"
	"net/http"
)

type VerifyHandlerDeps struct {
	*configs.Config
}

type VerifyHandler struct {
	*configs.Config
}

func NewVerifyHandler(router *http.ServeMux, deps VerifyHandlerDeps) {
	handler := &VerifyHandler{deps.Config}

	router.HandleFunc("POST /send", handler.Send())
	router.HandleFunc("GET /send/{hash}", handler.Verify())
}

func (handler *VerifyHandler) Send() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

func (handler *VerifyHandler) Verify() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}
