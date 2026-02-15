package errors

import (
	"errors"
)

// ItemNotFound - ошибка возвращаемая при ненайденной сущности
type ItemNotFound struct {
	Message string `json:"message"`
}

func (item *ItemNotFound) Error() string {
	return item.Message
}

func (item *ItemNotFound) Is(target error) bool {
	var itemNotFound *ItemNotFound
	ok := errors.As(target, &itemNotFound)
	return ok
}

func NewItemNotFound(message string) *ItemNotFound {
	return &ItemNotFound{Message: message}
}

var (
	ErrUserNotFound = errors.New("user not found")
)
