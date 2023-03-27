package pdf

import (
	"bytes"
	"fmt"
	"github.com/pablogolobaro/chequery/internal/domain/entity"
	"github.com/pablogolobaro/chequery/pkg/htmltopdf"
	"github.com/pablogolobaro/chequery/pkg/renderer"
	"path/filepath"
)

const dirPath string = "./media/pdf"
const checkTemplateName = "check"

type pdfStorage struct {
	t *renderer.Template
}

func NewPdfStorage(t *renderer.Template) *pdfStorage {
	return &pdfStorage{t: t}
}

func (p pdfStorage) GenerateCheckPDF(check entity.OrderCheck) (string, error) {
	filePath := filepath.Join(dirPath, fmt.Sprint(check.GetId(), ".pdf"))

	var b bytes.Buffer

	err := p.t.Render(&b, checkTemplateName, check, nil)
	if err != nil {
		return "", err
	}

	err = htmltopdf.GeneratePDF(filePath, &b)
	if err != nil {
		return "", err
	}

	return filePath, nil
}
