package utils

import (
	"bufio"
	"io"
	"os"
	"path/filepath"
)

func CreateFile(fileName, content string) error {
	dir := filepath.Dir(fileName)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}

	file, err := os.OpenFile(filepath.Join(fileName), os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	if _, err := file.WriteString(content); err != nil {
		return err
	}
	return nil
}

func ReadFile(path string) (string, error) {
	file, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		return "", err
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	var result string

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		result += string(line) + "\n"
	}

	return result, nil
}

func AddContentFile(path, content string) error {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	file.WriteString(content)
	return nil
}
