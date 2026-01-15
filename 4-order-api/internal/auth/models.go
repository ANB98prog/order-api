package auth

type AuthCode struct {
	SessionId string `json:"sessionId"`
	Code      string `json:"code"`
	Phone     string `json:"phone"`
}

type AuthCodes map[string]AuthCode

func (codes AuthCodes) Upsert(sessionId, code string, phone string) AuthCode {
	authCode := AuthCode{SessionId: sessionId, Code: code, Phone: phone}
	codes[sessionId] = authCode

	return authCode
}

func (codes AuthCodes) GetBySessionId(sessionId string) (AuthCode, bool) {
	code, ok := codes[sessionId]
	return code, ok
}

func (codes AuthCodes) Delete(sessionId string) {
	delete(codes, sessionId)
}
