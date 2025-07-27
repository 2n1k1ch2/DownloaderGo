package fetcher

import (
	"testing"
)

func TestIsAllowed(t *testing.T) {
	allowedFileURL := "https://www.w3.org/WAI/ER/tests/xhtml/testfiles/resources/pdf/dummy.pdf"
	notAllowedFileURL := "https://www.w3.org/WAI/ER/tests/xhtml/testfiles/resources/pdf/dummy.gps"
	if !IsAllowedFile(allowedFileURL) {
		t.Errorf("The allowed file URL %s is not allowed", allowedFileURL)
	}
	if IsAllowedFile(notAllowedFileURL) {
		t.Errorf("The notallowed file URL %s is allowed", allowedFileURL)
	}
}
func TestDownload(t *testing.T) {
	fileURL := "https://www.w3.org/WAI/ER/tests/xhtml/testfiles/resources/pdf/dummy.pdf"
	slb, err := DownloadFile(fileURL)
	if err != nil {
		t.Error(err)
	}
	if len(slb) == 0 {
		t.Errorf("The file URL %s is empty", fileURL)
	}

}
