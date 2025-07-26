package archiver

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func CreateZip(dest string, files []string) error {
	zipfile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer zipfile.Close()

	zipWriter := zip.NewWriter(zipfile)
	defer zipWriter.Close()

	for _, file := range files {
		fmt.Println("Archiving file:", file)

		f, err := os.Open(file)
		if err != nil {
			fmt.Println("Failed to open file:", file, err)
			continue
		}

		w, err := zipWriter.Create(filepath.Base(file))
		if err != nil {
			fmt.Println("Failed to create zip entry:", file, err)
			f.Close()
			continue
		}

		_, err = io.Copy(w, f)
		f.Close()
		if err != nil {
			fmt.Println("Failed to copy file into zip:", file, err)
			continue
		}
	}
	return nil
}
