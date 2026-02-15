package json

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"io"
)

func Decode[T any](body io.ReadCloser) (T, error) {
	var payload T
	err := json.NewDecoder(body).Decode(&payload)
	if err != nil {
		return payload, err
	}

	return payload, nil
}

func Encode[T any](writer io.WriteCloser, payload T) error {
	err := json.NewEncoder(writer).Encode(&payload)
	if err != nil {
		return err
	}

	return nil
}

func IsValid[T any](payload T) error {
	validator := validator.New()
	err := validator.Struct(payload)
	return err
}
