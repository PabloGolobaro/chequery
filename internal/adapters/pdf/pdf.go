package pdf

import (
	"bytes"
	"fmt"
	"github.com/pablogolobaro/chequery/internal/domain/entity"
	"github.com/pablogolobaro/chequery/pkg/htmltopdf"
	"html/template"
	"path/filepath"
)

const dirPath string = "./media/pdf"
const checkTemplateName = "checkbase"

type pdfStorage struct {
	template *template.Template
}

func NewPdfStorage(template *template.Template) *pdfStorage {
	return &pdfStorage{template: template}
}

func (p pdfStorage) GenerateCheckPDF(check entity.OrderCheck) (string, error) {
	filePath := filepath.Join(dirPath, fmt.Sprint(check.GetId(), ".pdf"))

	var b bytes.Buffer

	err := p.template.ExecuteTemplate(&b, checkTemplateName, check)
	if err != nil {
		return "", err
	}

	err = htmltopdf.GeneratePDF(filePath, &b)
	if err != nil {
		return "", err
	}

	return filePath, nil
}
