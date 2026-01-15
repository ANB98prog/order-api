package middlewares

import (
	"context"
	"github.com/ANB98prog/purple-school-homeworks/3-validation-api/pkg/response"
	"github.com/ANB98prog/purple-school-homeworks/4-order-api/configs"
	"github.com/ANB98prog/purple-school-homeworks/4-order-api/pkg/jwt"
	"net/http"
	"strings"
)

type CtxAuthorizationKey string

const (
	bearerPrefix                  = "Bearer "
	ErrInvalidAuthorizationMethod = "Invalid authorization method"
	ErrInvalidAuthorizationToken  = "Invalid authorization token"

	CtxUserSessionId CtxAuthorizationKey = "userSessionId"
	CtxUserPhone     CtxAuthorizationKey = "userPhone"
)

func Authorization(next http.Handler, config *configs.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorizationHeader := r.Header.Get("Authorization")
		if !strings.HasPrefix(bearerPrefix, authorizationHeader) {
			response.Unauthorized(w, response.ErrorMessage{Message: ErrInvalidAuthorizationMethod})
			return
		}

		bearerToken := strings.TrimPrefix("Bearer ", authorizationHeader)
		data, isValid := jwt.NewJWT(config.Auth.Secret).Parse(bearerToken)
		if !isValid {
			response.Unauthorized(w, response.ErrorMessage{Message: ErrInvalidAuthorizationToken})
			return
		}

		ctx := context.WithValue(r.Context(), CtxUserSessionId, data.SessionId)
		ctx = context.WithValue(ctx, CtxUserPhone, data.Phone)
		req := r.WithContext(ctx)

		next.ServeHTTP(w, req)
	})
}
