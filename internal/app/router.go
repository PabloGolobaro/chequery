package app

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const apiUri = "/api/v1"

func (a *Application) RegisterRouter() error {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/", "static")

	e.Renderer = a.renderer

	uiGroup := e.Group("")

	a.uiHandler.Register(uiGroup)

	apiGroup := uiGroup.Group(apiUri)

	a.checkHandler.Register(apiGroup)

	a.orderHandler.Register(apiGroup)

	a.healthHandler.Register(apiGroup)

	a.router = e

	return nil
}
