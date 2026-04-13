package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func GetPathSize(path string) (string, error) {
	info, err := os.Lstat(path)
	if err != nil {
		return "", err
	}

	var totalSize int64

	if !info.IsDir() {
		totalSize = info.Size()
	} else {
		entries, err := os.ReadDir(path)
		if err != nil {
			return "", err
		}

		for _, entry := range entries {
			entryPath := filepath.Join(path, entry.Name())

			entryInfo, err := os.Lstat(entryPath)
			if err != nil {
				return "", err
			}

			if entryInfo.IsDir() {
				continue
			}

			totalSize += entryInfo.Size()
		}
	}

	return fmt.Sprintf("%dB", totalSize), nil
}
