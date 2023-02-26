package pdf

import (
	"github.com/pablogolobaro/chequery/internal/domain/entity"
	"github.com/pablogolobaro/chequery/pkg/htmltopdf"
)

type pdfStorage struct {
}

func (p pdfStorage) GenerateCheckPDF(check entity.OrderCheck) error {
	htmltopdf.ExampleNewPDFGenerator()
	return nil
}
