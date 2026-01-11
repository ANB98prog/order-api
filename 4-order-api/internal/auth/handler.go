package auth

import (
	"fmt"
	"github.com/ANB98prog/purple-school-homeworks/3-validation-api/pkg/response"
	"github.com/ANB98prog/purple-school-homeworks/4-order-api/pkg/jwt"
	"github.com/ANB98prog/purple-school-homeworks/4-order-api/pkg/request"
	"log"
	"net/http"
)

type AuthHandler struct {
	authService *AuthCodeService
	jwt         *jwt.JWT
}

func NewAuthHandler(router *http.ServeMux, authService *AuthCodeService, jwt *jwt.JWT) {
	handler := &AuthHandler{
		authService: authService,
		jwt:         jwt,
	}

	router.HandleFunc("POST /auth/requestCode", handler.sendAuthCode())
	router.HandleFunc("POST /auth/verifyCode", handler.verifyAuthCode())
}

func (handler AuthHandler) sendAuthCode() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		payload, err := request.HandleBody[AuthCodeRequest](&w, r)
		if err != nil {
			return
		}

		authCode, err := handler.authService.GenerateAuthCode()
		if err != nil {
			response.InternalServerError(w, response.ErrorMessage{Message: err.Error()})
			return
		}

		// Номер телефона нужен только для отправки кода

		// Логируем сгенерированный токен сессии и код авторизации
		log.Printf("SessionId: %v Code: %v Phone: %s", authCode.SessionId, authCode.Code, payload.Phone)

		response.OK(w, AuthCodeResponse{SessionId: authCode.SessionId})
	}
}

func (handler AuthHandler) verifyAuthCode() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		payload, err := request.HandleBody[AuthCodeVerifyRequest](&w, r)
		if err != nil {
			return
		}

		isValid := handler.authService.VerifyAuthCode(payload.SessionId, payload.Code)
		if !isValid {
			response.BadRequest(w, response.ErrorMessage{Message: fmt.Errorf("invalid auth code").Error()})
			return
		}

		token, err := handler.jwt.Create(payload.SessionId)
		if err != nil {
			response.InternalServerError(w, response.ErrorMessage{Message: err.Error()})
			return
		}

		response.OK(w, AuthCodeVerifyResponse{Token: token})
	}
}
