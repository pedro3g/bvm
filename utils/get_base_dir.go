package utils

import (
	"errors"
	"os"
)

func GetBaseDir() (string, error) {
	baseDir := os.Getenv("BVM_DIR")

	if len(baseDir) == 0 {
		return "", errors.New("BVM base dir not found")
	}

	return baseDir, nil
}
