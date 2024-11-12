package main

import (
	"encoding/json"
	"fmt"
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
func WriteBytesToFile(path string, content []byte) error {
	return os.WriteFile(path, content, 0644)
}

func WriteTreeToFile(path string, root *Node) error {
	content, err := root.Json()
	if err != nil {
		return err
	}
	return WriteToFile(path, content)
}

func ReadTreeFromFile(path string) (*Node, error) {
	content, _ := ReadTextFile(path)

	var result Node

	err := json.Unmarshal([]byte(content), &result)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &result, nil
}
