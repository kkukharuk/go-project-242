package pathsize

import (
	"os"
	"path/filepath"
	"testing"
)

func TestGetSize_File(t *testing.T) {
	testFile := "testdata/test.txt"

	info, err := os.Stat(testFile)
	if err != nil {
		t.Fatalf("Failed to stat test file: %v", err)
	}

	size, err := GetSize(testFile, false, false)
	if err != nil {
		t.Fatalf("GetSize() error = %v", err)
	}

	expected := info.Size()
	if size != expected {
		t.Errorf("GetSize() = %d, want %d", size, expected)
	}
}

func TestGetSize_Directory(t *testing.T) {
	tmpDir := t.TempDir()

	file1 := filepath.Join(tmpDir, "file1.txt")
	file2 := filepath.Join(tmpDir, "file2.txt")

	err := os.WriteFile(file1, []byte("content1"), 0644)
	if err != nil {
		t.Fatalf("Failed to create file1: %v", err)
	}

	err = os.WriteFile(file2, []byte("content2"), 0644)
	if err != nil {
		t.Fatalf("Failed to create file2: %v", err)
	}

	size, err := GetSize(tmpDir, false, false)
	if err != nil {
		t.Fatalf("GetSize() error = %v", err)
	}

	expected := int64(8 + 8)
	if size != expected {
		t.Errorf("GetSize() = %d, want %d", size, expected)
	}
}

func TestGetSize_DirectoryWithSubdir(t *testing.T) {
	tmpDir := t.TempDir()

	file1 := filepath.Join(tmpDir, "file1.txt")
	subDir := filepath.Join(tmpDir, "subdir")
	file2 := filepath.Join(subDir, "file2.txt")

	err := os.WriteFile(file1, []byte("content1"), 0644)
	if err != nil {
		t.Fatalf("Failed to create file1: %v", err)
	}

	err = os.Mkdir(subDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create subdir: %v", err)
	}

	err = os.WriteFile(file2, []byte("content2"), 0644)
	if err != nil {
		t.Fatalf("Failed to create file2: %v", err)
	}

	size, err := GetSize(tmpDir, false, false)
	if err != nil {
		t.Fatalf("GetSize() error = %v", err)
	}

	expected := int64(8)
	if size != expected {
		t.Errorf("GetSize() non-recursive = %d, want %d (only first level files)", size, expected)
	}

	sizeRecursive, err := GetSize(tmpDir, false, true)
	if err != nil {
		t.Fatalf("GetSize() recursive error = %v", err)
	}

	expectedRecursive := int64(8 + 8)
	if sizeRecursive != expectedRecursive {
		t.Errorf("GetSize() recursive = %d, want %d", sizeRecursive, expectedRecursive)
	}
}

func TestGetSize_NotFound(t *testing.T) {
	_, err := GetSize("/nonexistent/path", false, false)
	if err == nil {
		t.Error("GetSize() expected error for nonexistent path, got nil")
	}
}

func TestGetSize_WithHiddenFiles(t *testing.T) {
	tmpDir := t.TempDir()

	visibleFile := filepath.Join(tmpDir, "visible.txt")
	hiddenFile := filepath.Join(tmpDir, ".hidden")

	err := os.WriteFile(visibleFile, []byte("visible"), 0644)
	if err != nil {
		t.Fatalf("Failed to create visible file: %v", err)
	}

	err = os.WriteFile(hiddenFile, []byte("hidden"), 0644)
	if err != nil {
		t.Fatalf("Failed to create hidden file: %v", err)
	}

	sizeWithoutHidden, err := GetSize(tmpDir, false, false)
	if err != nil {
		t.Fatalf("GetSize() error = %v", err)
	}

	sizeWithHidden, err := GetSize(tmpDir, true, false)
	if err != nil {
		t.Fatalf("GetSize() error = %v", err)
	}

	if sizeWithoutHidden != 7 {
		t.Errorf("GetSize() without --all = %d, want 7", sizeWithoutHidden)
	}

	if sizeWithHidden != 13 {
		t.Errorf("GetSize() with --all = %d, want 13", sizeWithHidden)
	}
}

func TestFormatSize(t *testing.T) {
	tests := []struct {
		name     string
		size     int64
		expected string
	}{
		{"bytes", 64, "64 B"},
		{"kilobytes", 1024, "1.00 kB"},
		{"kilobytes_fractional", 1536, "1.50 kB"},
		{"megabytes", 1048576, "1.00 MB"},
		{"megabytes_fractional", 1572864, "1.50 MB"},
		{"gigabytes", 1073741824, "1.00 GB"},
		{"gigabytes_fractional", 1610612736, "1.50 GB"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FormatSize(tt.size)
			if result != tt.expected {
				t.Errorf("FormatSize(%d) = %s, want %s", tt.size, result, tt.expected)
			}
		})
	}
}
