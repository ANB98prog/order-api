package service

import (
	"errors"
	"github.com/ANB98prog/order-api/internal/repository"
	helper "github.com/ANB98prog/order-api/pkg/helpers/auth"
	"time"
)

const (
	minTokenLen = 6
	sessionTtl  = time.Minute * 5
)

type AuthCodeService interface {
	GenerateAuthCode(userId uint) (AuthCode, error)
	GetAuthCode(sessionId string) (AuthCode, bool)
	DeleteAuthCode(sessionId string)
}

type authCodeService struct {
	authCodeRepo repository.AuthCodeRepository
}

func NewAuthCodeService(repo repository.AuthCodeRepository) AuthCodeService {
	return &authCodeService{authCodeRepo: repo}
}

var _ AuthCodeService = (*authCodeService)(nil)

func (service *authCodeService) GenerateAuthCode(userId uint) (AuthCode, error) {
	authCode, err := service.generateAuthCode(userId)
	if err != nil {
		return AuthCode{}, err
	}

	return authCode, nil
}

func (service *authCodeService) generateAuthCode(userId uint) (AuthCode, error) {
	authCode := repository.AuthCode{
		UserId: userId,
		Code:   helper.GenerateAuthCode(),
	}

	minSessionIdLen := minTokenLen
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			sessionId := helper.GenerateSessionId(minSessionIdLen)

			authCode.SessionId = sessionId

			ok, err := service.authCodeRepo.Save(authCode, sessionTtl)
			if err != nil {
				return AuthCode{}, err
			}

			if ok {
				return AuthCode{SessionId: authCode.SessionId, Code: authCode.Code, UserId: authCode.UserId}, nil
			}
		}
		// Если за n попыток не удалось сгенерировать уникальную сессию, то увеличиваем размер токена
		minSessionIdLen++
	}

	return AuthCode{}, errors.New("cannot create unique session")
}

func (service *authCodeService) GetAuthCode(sessionId string) (AuthCode, bool) {
	authCode, ok := service.authCodeRepo.GetBySessionId(sessionId)
	if !ok {
		return AuthCode{}, false
	}

	return AuthCode{SessionId: authCode.SessionId, UserId: authCode.UserId, Code: authCode.Code}, ok
}

func (service *authCodeService) DeleteAuthCode(sessionId string) {
	service.authCodeRepo.Delete(sessionId)
}
