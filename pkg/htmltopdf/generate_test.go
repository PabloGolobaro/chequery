package htmltopdf

import (
	"bytes"
	"encoding/json"
	"github.com/pablogolobaro/chequery/internal/domain/entity"
	"github.com/pablogolobaro/chequery/pkg/templ"
	"path/filepath"
	"strconv"
	"testing"
)

func testOrder() entity.OrderCheck {
	m := map[string]interface{}{"vegetables": 1, "pork": 3}
	indent, err := json.MarshalIndent(m, "", "\t")
	if err != nil {
		return entity.OrderCheck{}
	}
	orderCheck := entity.NewCheckBuilder().SetId(1).SetCheckType(entity.Kitchen).SetStatus("created").SetPrinterId("123").SetOrder(string(indent)).Build()
	return orderCheck
}

func TestGeneratePDF(t *testing.T) {
	temp, err := templ.ParseTemplate("./testdata")
	if err != nil {
		t.Error(err)
		return
	}

	var b bytes.Buffer

	orderCheck := testOrder()
	m := map[string]interface{}{
		"id":        orderCheck.Id(),
		"checkType": orderCheck.CheckType(),
		"printerId": orderCheck.PrinterId(),
		"order":     orderCheck.Order(),
	}

	err = temp.ExecuteTemplate(&b, "base", m)
	if err != nil {
		t.Error(err)
		return
	}

	dirPath := "./testdata/pdf"
	filePath := filepath.Join(dirPath, strconv.Itoa(orderCheck.Id())+".pdf")

	type args struct {
		path   string
		buffer *bytes.Buffer
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{{name: "Simple", args: args{path: filePath, buffer: &b}, wantErr: false}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := GeneratePDF(tt.args.path, tt.args.buffer); (err != nil) != tt.wantErr {
				t.Errorf("GeneratePDF() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
