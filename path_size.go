package code

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func GetPathSize(path string, all, recursive, human bool) (string, error) {
	size, err := getPathSizeBytes(path, all, recursive)
	if err != nil {
		return "", err
	}

	return formatSize(size, human), nil
}

func getPathSizeBytes(path string, all, recursive bool) (int64, error) {
	info, err := os.Lstat(path)
	if err != nil {
		return 0, err
	}

	if !info.IsDir() {
		return info.Size(), nil
	}

	entries, err := os.ReadDir(path)
	if err != nil {
		return 0, err
	}

	var totalSize int64

	for _, entry := range entries {
		if !all && strings.HasPrefix(entry.Name(), ".") {
			continue
		}

		entryPath := filepath.Join(path, entry.Name())
		entryInfo, err := os.Lstat(entryPath)
		if err != nil {
			return 0, err
		}

		if entryInfo.IsDir() {
			if recursive {
				dirSize, err := getPathSizeBytes(entryPath, all, recursive)
				if err != nil {
					return 0, err
				}
				totalSize += dirSize
			}
			continue
		}

		totalSize += entryInfo.Size()
	}

	return totalSize, nil
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
