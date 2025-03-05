package utils

import (
	"fmt"
	"os"
)

// CreateDir ensures the output directory exists
func CreateDir(dir string) error {
	if checkDirExists(dir) {
		count := 1
		for checkDirExists(fmt.Sprintf("%s_%d", dir, count)) {
			count++
		}
		dir = fmt.Sprintf("%s_%d", dir, count)
	}
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}
	return nil
}

func checkDirExists(dir string) bool {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return false
	}
	return true
}
