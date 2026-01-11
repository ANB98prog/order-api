package auth

import (
	"github.com/ANB98prog/purple-school-homeworks/4-order-api/pkg/helpers/auth"
	"github.com/ANB98prog/purple-school-homeworks/4-order-api/pkg/helpers/file"
)

const codesFileName = "authCodes.json"

type AuthCodeService struct {
}

func NewAuthService() *AuthCodeService {
	return new(AuthCodeService)
}

func (service *AuthCodeService) GenerateAuthCode() (AuthCode, error) {

	// Генерируем код и сессию
	code := auth.GenerateAuthCode()
	sessionId := auth.GenerateSessionId()

	// Читаем из файла коды
	authCodes, err := getCodesFromFile()
	if err != nil {
		return AuthCode{}, err
	}

	// Добавляем или обновляем код для пользователя
	authCodes.Upsert(sessionId, code)

	// Сохраняем коды
	err = file.WriteFile(codesFileName, &authCodes)
	if err != nil {
		return AuthCode{}, err
	}

	// Возвращаем код
	return AuthCode{Code: code, SessionId: sessionId}, nil
}

func (service *AuthCodeService) VerifyAuthCode(sessionId, code string) bool {
	// Достаем коды
	authCodes, err := getCodesFromFile()
	if err != nil {
		return false
	}

	// Проверяем на соответствие
	authSession, ok := authCodes.GetBySessionId(sessionId)
	if !ok || authSession.Code != code {
		return false
	}

	return true
}

func getCodesFromFile() (*AuthCodes, error) {
	authCodes, err := file.ReadFile[AuthCodes](codesFileName)

	if err != nil {
		return nil, err
	}

	if authCodes == nil {
		return &AuthCodes{}, nil
	}

	return authCodes, nil
}
