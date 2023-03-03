package templ

import (
	"html/template"
	"path/filepath"
)

func ParseTemplate(templatePath string) (*template.Template, error) {
	pattern := filepath.Join(templatePath, "*.html")

	glob, err := filepath.Glob(pattern)
	if err != nil {
		return nil, err
	}

	t, err := template.ParseFiles(glob...)
	if err != nil {
		return nil, err
	}

	return t, nil
}
