package main

import "html/template"

func createTemplate(path string) *template.Template {
	t := template.Must(template.ParseFiles("./static/templates/base.html", "./static/templates/check.html"))
	return t
}
