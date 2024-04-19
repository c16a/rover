package utils

import (
	"encoding/json"
	"os"
)

func ParseJsonFile[T any](path string) (*T, error) {
	var config T
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
