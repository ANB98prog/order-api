package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"strconv"
)

const (
	SessionIdKey = "sessionId"
	UserPhoneKey = "phone"
	UserIdKey    = "userId"
)

// JWT - содержит данные для создания JWT токена
type JWT struct {
	secret string
}

type JWTData struct {
	SessionId string
	Phone     string
	UserId    uint
}

// NewJWT - создает новый экземпляр JWT
func NewJWT(secret string) *JWT {
	return &JWT{secret}
}

// Create - Создает jwt токен
func (j *JWT) Create(sessionId, phone string, userId uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		SessionIdKey: sessionId,
		UserPhoneKey: phone,
		UserIdKey:    strconv.Itoa(int(userId)),
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
	sessionId, ok := GetClaim(SessionIdKey, token)
	if !ok {
		return nil, false
	}
	phone, ok := GetClaim(UserPhoneKey, token)
	if !ok {
		return nil, false
	}
	userIdStr, ok := GetClaim(UserIdKey, token)
	if !ok {
		return nil, false
	}
	userId, err := strconv.ParseUint(userIdStr, 10, 32)
	if err != nil {
		return nil, false
	}

	return &JWTData{
		SessionId: sessionId,
		Phone:     phone,
		UserId:    uint(userId),
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
