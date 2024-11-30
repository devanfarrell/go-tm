package io

import (
	"os"
)

func ReadFile(filePath string) (string, error) {
	dat, err := os.ReadFile(filePath)
	return string(dat), err
}
