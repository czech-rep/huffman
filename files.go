package main

import (
	"encoding/json"
	"os"
)

func ReadTextFile(path string) (string, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func WriteToFile(path string, content string) error {
	data := []byte{}
	for _, char := range content {
		data = append(data, byte(char))
	}
	return os.WriteFile(path, data, 0644)
}

func WriteTreeToFile(path string, root *Node) error {
	content, err := root.Json()
	if err != nil {
		return err
	}
	return WriteToFile(path, content)
}

func ReadTreeFromFile(path string) (*Node, error) {
	content, _ := os.ReadFile(path)

	var result Node

	err := json.Unmarshal(content, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
