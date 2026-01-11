package json

import (
	"bytes"
	"encoding/json"
	"fmt"
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

func DecodeBytes[T any](data []byte) (*T, error) {
	var payload T
	dec := json.NewDecoder(bytes.NewReader(data))
	dec.DisallowUnknownFields()

	if err := dec.Decode(&payload); err != nil {
		return nil, fmt.Errorf("decode JSON: %w", err)
	}

	return &payload, nil
}

func Encode[T any](writer io.WriteCloser, payload T) error {
	err := json.NewEncoder(writer).Encode(&payload)
	if err != nil {
		return err
	}

	return nil
}

func IsValid[T any](payload T) error {
	v := validator.New()
	err := v.Struct(payload)
	return err
}
