package archiver

import (
	"archive/zip"
	"fmt"
	"os"
	"testing"
)

func TestCreateZip(t *testing.T) {
	// Подготовка временных файлов
	expected := make(map[string]bool)
	filePaths := []string{}

	for i := 0; i < 3; i++ {
		name := fmt.Sprintf("Test%d.txt", i)
		err := os.WriteFile(name, []byte(fmt.Sprintf("content %d", i)), 0644)
		if err != nil {
			t.Fatalf("failed to create test file: %v", err)
		}
		expected[name] = false
		filePaths = append(filePaths, name)
	}

	archivePath := "test.zip"
	defer func() {
		_ = os.Remove(archivePath)
		for _, f := range filePaths {
			_ = os.Remove(f)
		}
	}()

	if err := CreateZip(archivePath, filePaths); err != nil {
		t.Fatalf("CreateZip failed: %v", err)
	}

	r, err := zip.OpenReader(archivePath)
	if err != nil {
		t.Fatalf("Failed to open archive: %v", err)
	}
	defer r.Close()

	for _, f := range r.File {
		if _, ok := expected[f.Name]; ok {
			expected[f.Name] = true
		}
	}

	for name, found := range expected {
		if !found {
			t.Errorf("Expected file %s not found in archive", name)
		}
	}
}
