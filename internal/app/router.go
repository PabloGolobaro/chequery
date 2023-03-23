package app

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"html/template"
	"io"
)

const apiUri = "/api/v1"

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func (a *Application) RegisterRouter() error {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/", "static")

	t := &Template{
		templates: a.t,
	}

	e.Renderer = t

	uiGroup := e.Group("/ui")

	a.uiHandler.Register(uiGroup)

	apiGroup := e.Group(apiUri)

	a.checkHandler.Register(apiGroup)

	a.orderHandler.Register(apiGroup)

	a.healthHandler.Register(apiGroup)

	a.router = e

	return nil
}
