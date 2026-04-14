package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func GetPathSize(path string, human, all bool) (string, error) {
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
			if !all && strings.HasPrefix(entry.Name(), ".") {
				continue
			}

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

	return formatSize(totalSize, human), nil
}

func formatSize(size int64, human bool) string {
	if !human {
		return fmt.Sprintf("%dB", size)
	}
	units := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}
	value := float64(size)
	unitIndex := 0

	for value >= 1024 && unitIndex < len(units)-1 {
		value /= 1024
		unitIndex++
	}

	if unitIndex == 0 {
		return fmt.Sprintf("%dB", size)
	}

	return fmt.Sprintf("%.1f%s", value, units[unitIndex])
}
