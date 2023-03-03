package templ

import (
	"testing"
)

func TestParseTemplate(t *testing.T) {
	tests := []struct {
		name    string
		path    string
		wantErr bool
	}{
		{name: "Simple", path: "./testdata/static/templates", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ParseTemplate(tt.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseTemplate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
