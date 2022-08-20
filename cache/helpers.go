package cache

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
