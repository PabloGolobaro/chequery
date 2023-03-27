package htmltopdf

import (
	"bytes"
	"github.com/pablogolobaro/chequery/internal/domain/entity"
	"github.com/pablogolobaro/chequery/pkg/renderer"
	"path/filepath"
	"strconv"
	"testing"
)

const (
	checkTemplateName = "check"
	templateDir       = "./testdata/templates"
	pdfFilesPath      = "./testdata/pdf"
	testPath          = "/testdata"
)

func testOrder() entity.OrderCheck {
	orderData := entity.Order{
		PointID: 1,
		Products: []entity.Product{
			{Name: "Meat", Quantity: 3, Price: 145},
			{Name: "vegetables", Quantity: 2, Price: 32},
			{Name: "Juice", Quantity: 1, Price: 48}},
	}
	orderCheck := entity.NewCheckBuilder().SetId(1).SetCheckType(entity.Kitchen).SetStatus("created").SetPrinterId("123").SetOrder(orderData.Details()).Build()
	return orderCheck
}

func TestGeneratePDF(t *testing.T) {
	renderer := renderer.New()

	err := renderer.LoadTemplates(templateDir)
	if err != nil {
		t.Log(err)
		return
	}
	var b bytes.Buffer

	orderCheck := testOrder()

	err = renderer.Render(&b, checkTemplateName, orderCheck, nil)
	if err != nil {
		t.Log(err)
		return
	}

	err = FindWKHTMLTOPDF(testPath)
	if err != nil {
		t.Fatal(err)
	}

	filePath := filepath.Join(pdfFilesPath, strconv.Itoa(orderCheck.GetId())+".pdf")

	type args struct {
		path   string
		buffer *bytes.Buffer
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{{
		name: "Simple", args: args{path: filePath, buffer: &b}, wantErr: false,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := GeneratePDF(tt.args.path, tt.args.buffer); (err != nil) != tt.wantErr {
				t.Errorf("GeneratePDF() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
