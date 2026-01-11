package jwt

import "github.com/golang-jwt/jwt/v5"

// JWT - содержит данные для создания JWT токена
type JWT struct {
	secret string
}

// NewJWT - создает новый экземпляр JWT
func NewJWT(secret string) *JWT {
	return &JWT{secret}
}

// Create - Создает jwt токен
func (j *JWT) Create(sessionId string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sessionId": sessionId,
	})

	s, err := token.SignedString([]byte(j.secret))
	if err != nil {
		return "", err
	}
	return s, nil
}
