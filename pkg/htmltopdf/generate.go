package htmltopdf

import (
	"bytes"
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

func GeneratePDF(path string, buffer *bytes.Buffer) error {
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		return err
	}

	pdfg.Dpi.Set(300)
	pdfg.Orientation.Set(wkhtmltopdf.OrientationLandscape)
	pdfg.Grayscale.Set(true)

	pageReader := wkhtmltopdf.NewPageReader(buffer)

	pageReader.FooterRight.Set("[page]")
	pageReader.FooterFontSize.Set(10)
	pageReader.Zoom.Set(0.95)

	pdfg.AddPage(pageReader)

	err = pdfg.Create()
	if err != nil {
		return err
	}

	// Write buffer contents to file on disk
	err = pdfg.WriteFile(path)
	if err != nil {
		return err
	}

	return nil
}
