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

	group := e.Group(apiUri)

	a.checkHandler.Register(group)

	a.orderHandler.Register(group)

	a.healthHandler.Register(group)

	a.router = e

	return nil
}
