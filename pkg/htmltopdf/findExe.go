package htmltopdf

import (
	"os"
	"path/filepath"
)

func FindWKHTMLTOPDF(exePath string) error {
	exPath, err := os.Getwd()
	if err != nil {
		return err
	}

	toolpath := filepath.Join(exPath, exePath)

	err = os.Setenv("WKHTMLTOPDF_PATH", toolpath)
	if err != nil {
		return err
	}
	return nil
}
