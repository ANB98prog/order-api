package jwt

import "github.com/golang-jwt/jwt/v5"

// JWT - содержит данные для создания JWT токена
type JWT struct {
	secret string
}

type JWTData struct {
	SessionId string
	Phone     string
}

// NewJWT - создает новый экземпляр JWT
func NewJWT(secret string) *JWT {
	return &JWT{secret}
}

// Create - Создает jwt токен
func (j *JWT) Create(sessionId, phone string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sessionId": sessionId,
		"phone":     phone,
	})

	s, err := token.SignedString([]byte(j.secret))
	if err != nil {
		return "", err
	}
	return s, nil
}

// Parse - парсит данные из токена и проверяет его валидность
func (j *JWT) Parse(tokenString string) (*JWTData, bool) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.secret), nil
	})
	if err != nil {
		return nil, false
	}
	sessionId, ok := GetClaim("sessionId", token)
	if !ok {
		return nil, false
	}
	phone, ok := GetClaim("phone", token)
	if !ok {
		return nil, false
	}
	return &JWTData{
		SessionId: sessionId,
		Phone:     phone,
	}, token.Valid
}

func GetClaim(key string, token *jwt.Token) (string, bool) {
	value, ok := token.Claims.(jwt.MapClaims)[key]
	if !ok {
		return "", false
	}
	valueOfStr, ok := value.(string)
	return valueOfStr, ok
}
