package utils

import (
	"io"
	"os"
	"path/filepath"
)

const (
	FILE_NAME = "togos.json"
)

func GetFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, FILE_NAME), nil
}

func ensureFileExistsAndInit() error {
	filePath, err := GetFilePath()
	if err != nil {
		return err
	}

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
		_, err = file.Write([]byte("{\"tasks\":[]}"))
		file.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

func ReadFile() ([]byte, error) {
	err := ensureFileExistsAndInit()
	if err != nil {
		return nil, err
	}
	filePath, err := GetFilePath()
	if err != nil {
		return nil, err
	}
	file, err := os.OpenFile(filePath, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return content, nil
}

func WriteFile(content string) error {
	err := ensureFileExistsAndInit()
	if err != nil {
		return err
	}
	filePath, err := GetFilePath()
	if err != nil {
		return err
	}
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(content)
	if err != nil {
		return err
	}
	return nil
}
