package verify

import (
	"errors"
	"github.com/ANB98prog/purple-school-homeworks/3-validation-api/pkg/request"
	"github.com/ANB98prog/purple-school-homeworks/3-validation-api/pkg/response"
	"net/http"
)

type VerifyHandlerDeps struct {
	*configs.Config
	VerificationService VerificationService
}

type VerifyHandler struct {
	config              *configs.Config
	verificationService VerificationService
}

func NewVerifyHandler(router *http.ServeMux, deps VerifyHandlerDeps) {
	handler := &VerifyHandler{
		config:              deps.Config,
		verificationService: deps.VerificationService,
	}

	router.HandleFunc("POST /send", handler.Send())
	router.HandleFunc("GET /verify/{hash}", handler.Verify())
}

func (handler *VerifyHandler) Send() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		payload, err := request.HandleBody[EmailVerification](&w, r)
		if err != nil {
			return
		}

		err = handler.verificationService.SendVerification(payload.Email)
		if err != nil {
			response.InternalServerError(w, response.ErrorMessage{Message: err.Error()})
			return
		}

		response.OKWithData(w, map[string]any{})
		return
	}
}

func (handler *VerifyHandler) Verify() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hash := r.PathValue("hash")
		if hash == "" {
			response.BadRequest(w, errors.New("no hash provided in link"))
			return
		}

		if isVerified := handler.verificationService.Verify(hash); isVerified {
			response.OKWithData(w, map[string]any{})
			return
		}

		response.NotFound(w, response.ErrorMessage{Message: "Email verification is not found"})
	}
}
