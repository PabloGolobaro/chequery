package renderer

import (
	"os"
	"testing"
)

func TestTemplate_LoadTemplates(t1 *testing.T) {

	tests := []struct {
		name            string
		baseTemplateDir string
		wantErr         bool
	}{{
		name:            "Simple",
		baseTemplateDir: "./testdata/templates",
		wantErr:         false,
	}}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := New()
			if err := t.LoadTemplates(tt.baseTemplateDir); (err != nil) != tt.wantErr {
				t1.Errorf("LoadTemplates() error = %v, wantErr %v", err, tt.wantErr)
			}
			for s, template := range t.templatesMap {
				t1.Log(s)
				t1.Log(template.Name())
				t1.Log(template.DefinedTemplates())
				err := template.ExecuteTemplate(os.Stdout, "base", nil)
				if err != nil {
					t1.Log(err)
					return
				}
			}

		})
	}
}
