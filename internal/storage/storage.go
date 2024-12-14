package storage

import (
	"encoding/json"
	"os"
)

// declare a storage type
type Storage[T any] struct {
	Filename string
}

// creates a new storage obj
func NewStorage[T any](filename string) *Storage[T] {
	return &Storage[T]{Filename: filename}
}

func (s *Storage[T]) Save(data T) error {
	// save data to json
	fileData, err := json.MarshalIndent(data, "", "    ")

	if err != nil {
		return err
	}

	// 0644 = file permission (rw- r-- r--)
	return os.WriteFile(s.Filename, fileData, 0644)
}

func (s *Storage[T]) Load(data *T) error {
	fileData, err := os.ReadFile(s.Filename)

	if err != nil {
		return err
	}
	// converts json data and populate the data param
	return json.Unmarshal(fileData, data)
}