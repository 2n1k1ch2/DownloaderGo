package fetcher

import (
	"fmt"
	"io"
	"net/http"
	"path/filepath"
	"strings"
)

func DownloadFile(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bad status: %s", resp.Status)
	}

	return io.ReadAll(resp.Body)
}

func IsAllowedFile(url string) bool {
	ext := strings.ToLower(filepath.Ext(url))
	return ext == ".pdf" || ext == ".jpeg" || ext == ".jpg"
}
