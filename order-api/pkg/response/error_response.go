package response

import "encoding/json"

type ErrorMessage struct {
	Message string `json:"message"`
}

func NewErrorMessage(err error) *ErrorMessage {
	return &ErrorMessage{Message: err.Error()}
}

func (message *ErrorMessage) Error() string {
	data, err := json.Marshal(message)
	if err != nil {
		return message.Message
	}
	return string(data)
}
