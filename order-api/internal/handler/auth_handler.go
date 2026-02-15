package handler

import (
	goerrors "errors"
	"fmt"
	"github.com/ANB98prog/order-api/internal/service"
	"github.com/ANB98prog/order-api/pkg/errors"
	"github.com/ANB98prog/order-api/pkg/jwt"
	"github.com/ANB98prog/order-api/pkg/request"
	"github.com/ANB98prog/order-api/pkg/response"
	"log"
	"net/http"
)

type AuthHandler struct {
	userService service.UserService
	authService service.AuthCodeService
	jwt         *jwt.JWT
}

type AuthHandlerDeps struct {
	service.UserService
	service.AuthCodeService
	*jwt.JWT
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		userService: deps.UserService,
		authService: deps.AuthCodeService,
		jwt:         deps.JWT,
	}

	router.HandleFunc("POST /auth/requestCode", handler.requestAuthCode())
	router.HandleFunc("POST /auth/verifyCode", handler.verifyAuthCode())
}

func (h AuthHandler) requestAuthCode() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		payload, err := request.HandleBody[AuthCodeRequest](&w, r)
		if err != nil {
			return
		}

		user, err := h.getOrCreateUser(payload.Phone)
		if err != nil {
			response.InternalServerError(w, response.ErrorMessage{Message: err.Error()})
			return
		}

		authCode, err := h.authService.GenerateAuthCode(user.Id)
		if err != nil {
			response.InternalServerError(w, response.ErrorMessage{Message: err.Error()})
			return
		}

		// Логируем сгенерированный токен сессии и код авторизации
		log.Printf("SessionId: %v Code: %v Phone: %s", authCode.SessionId, authCode.Code, payload.Phone)

		response.OKWithData(w, AuthCodeResponse{SessionId: authCode.SessionId})
	}
}

func (h AuthHandler) getOrCreateUser(phone string) (*service.User, error) {
	// проверяем существование пользователя, если не существует создаем нового
	user, err := h.userService.GetByPhone(phone)
	if err != nil {
		if goerrors.Is(err, errors.ErrUserNotFound) {
			user, err = h.userService.Create(phone)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	return user, nil
}

func (h AuthHandler) verifyAuthCode() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		payload, err := request.HandleBody[AuthCodeVerifyRequest](&w, r)
		if err != nil {
			return
		}

		authCode, ok := h.authService.GetAuthCode(payload.SessionId)
		if !ok || authCode.Code != payload.Code {
			response.Unauthorized(w, response.ErrorMessage{Message: fmt.Errorf("invalid auth code").Error()})
			return
		}

		user, err := h.userService.GetById(authCode.UserId)
		if err != nil {
			if goerrors.Is(err, errors.ErrUserNotFound) {
				response.Unauthorized(w, response.ErrorMessage{Message: err.Error()})
				return
			}

			response.InternalServerError(w, response.ErrorMessage{Message: err.Error()})
			return
		}

		token, err := h.jwt.Create(payload.SessionId, user.Phone, user.Id)
		if err != nil {
			response.InternalServerError(w, response.ErrorMessage{Message: err.Error()})
			return
		}

		response.OKWithData(w, AuthCodeVerifyResponse{Token: token})
	}
}
