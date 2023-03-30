package renderer

import (
	"github.com/labstack/echo/v4"
	"html/template"
	"io"
	"log"
	"path/filepath"
	"strings"
)

const baseTemplateName = "base"
const snippetKey = "snippet"

type Template struct {
	templatesMap map[string]*template.Template
}

func New() *Template {
	return &Template{templatesMap: make(map[string]*template.Template)}
}

func (t *Template) LoadTemplates(baseTemplateDir string) error {
	contentTemplates, err := filepath.Glob(baseTemplateDir + "/*/*.html")
	if err != nil {
		return err
	}

	templateBase := template.Must(template.ParseGlob(baseTemplateDir + "/*.html"))

	snippetTemplates := template.Must(template.ParseGlob(baseTemplateDir + "/snippets/*.html"))

	t.templatesMap[snippetKey] = snippetTemplates

	for _, contentTemplate := range contentTemplates {
		if strings.Contains(contentTemplate, snippetKey) {
			continue
		}

		templateName := strings.Split(filepath.Base(contentTemplate), ".")[0]

		templateBaseClone, err := templateBase.Clone()
		if err != nil {
			return err
		}

		resultTemplate, err := templateBaseClone.ParseFiles(contentTemplate)
		if err != nil {
			return err
		}

		t.templatesMap[templateName] = resultTemplate
	}

	t.logTemplates()

	return nil
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	if strings.Contains(name, snippetKey) {
		return t.templatesMap[snippetKey].ExecuteTemplate(w, name, data)
	}
	return t.templatesMap[name].ExecuteTemplate(w, baseTemplateName, data)
}

// Only for debug purposes
func (t *Template) logTemplates() {
	for name, t2 := range t.templatesMap {
		log.Println(t2.DefinedTemplates())
		log.Println(t2.Name())
		log.Println(name)
	}
}
