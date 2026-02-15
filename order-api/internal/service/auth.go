package service

type AuthCode struct {
	SessionId string
	UserId    uint
	Code      string
}
