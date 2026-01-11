package auth

type AuthCode struct {
	SessionId string `json:"sessionId"`
	Code      string `json:"code"`
}

type AuthCodes map[string]AuthCode

func (codes AuthCodes) Upsert(sessionId, code string) {
	codes[sessionId] = AuthCode{SessionId: sessionId, Code: code}
}

func (codes AuthCodes) GetBySessionId(sessionId string) (AuthCode, bool) {
	code, ok := codes[sessionId]
	return code, ok
}
