package auth

type AuthCodeRequest struct {
	Phone string `json:"phone" validate:"required"`
}

type AuthCodeResponse struct {
	SessionId string `json:"sessionId"`
}

type AuthCodeVerifyRequest struct {
	SessionId string `json:"sessionId" validate:"required"`
	Code      string `json:"code" validate:"required"`
}

type AuthCodeVerifyResponse struct {
	Token string `json:"token"`
}
