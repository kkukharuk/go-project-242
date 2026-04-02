package pathsize

import (
	"fmt"
	"os"
	"path/filepath"
)

func GetSize(path string, includeHidden bool, recursive bool) (int64, error) {
	info, err := os.Stat(path)
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

	var total int64
	for _, entry := range entries {
		if !includeHidden && isHidden(entry.Name()) {
			continue
		}
		info, err := entry.Info()
		if err != nil {
			return 0, err
		}
		if info.IsDir() {
			if recursive {
				subPath := filepath.Join(path, entry.Name())
				subSize, err := GetSize(subPath, includeHidden, recursive)
				if err != nil {
					return 0, err
				}
				total += subSize
			}
		} else {
			total += info.Size()
		}
	}

	return total, nil
}

func isHidden(name string) bool {
	return len(name) > 0 && name[0] == '.'
}

func FormatSize(size int64) string {
	const (
		KB = 1024
		MB = KB * 1024
		GB = MB * 1024
	)

	switch {
	case size >= GB:
		return fmt.Sprintf("%.2f GB", float64(size)/GB)
	case size >= MB:
		return fmt.Sprintf("%.2f MB", float64(size)/MB)
	case size >= KB:
		return fmt.Sprintf("%.2f kB", float64(size)/KB)
	default:
		return fmt.Sprintf("%d B", size)
	}
}
