package internal

import (
	"os"
)

func readFromFile(fileName string) (string, error) {
	htmlData, err := os.ReadFile(fileName)
	if err != nil {
		return "", err
	}
	return string(htmlData), nil
}

func CleanUp(fileName string) error {
	err := os.Remove(fileName)
	return err
}
