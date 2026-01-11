package file

import (
	"fmt"
	"github.com/ANB98prog/purple-school-homeworks/3-validation-api/pkg/json"
	"os"
)

func ReadFile[T any](filePath string) (*T, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("read file %s: %w", filePath, err)
	}

	if len(data) == 0 {
		return nil, nil
	}

	// Валидация JSON
	fileData, err := json.DecodeBytes[T](data)
	if err != nil {
		return fileData, err
	}

	return fileData, nil
}

func WriteFile[T any](filePath string, data *T) error {
	file, err := os.OpenFile(filePath, os.O_CREATE, os.ModePerm)
	if err != nil {
		return fmt.Errorf("could not write to file %s. Error: %e", filePath, err)
	}
	defer file.Close()

	err = json.Encode(file, data)
	if err != nil {
		return fmt.Errorf("could not write to file %s. Error: %e", filePath, err)
	}

	return nil
}
