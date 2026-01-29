package auth

import "math/rand/v2"

const intCharset = "0123456789"
const alphaNumericCharset = "abcdefghijklmnopqrstuvwxyzZYXWVUTSRQPONMLKJIHGFEDCBA!@#$%^&*_-=+0123456789"

// GenerateSessionId - генерирует токен сессии
func GenerateSessionId(minLen int) string {
	token := make([]byte, minLen)
	for i := range token {
		token[i] = alphaNumericCharset[rand.IntN(len(alphaNumericCharset))]
	}

	return string(token)
}

// GenerateAuthCode - генерирует 4-х значный код авторизации
func GenerateAuthCode() string {
	code := make([]byte, 4)
	for i := range code {
		code[i] = intCharset[rand.IntN(len(intCharset))]
	}

	return string(code)
}
