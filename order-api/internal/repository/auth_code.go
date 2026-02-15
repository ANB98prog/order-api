package repository

type AuthCode struct {
	SessionId string
	UserId    uint
	Code      string
}
