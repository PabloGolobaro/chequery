package app

import "github.com/labstack/echo/v4"

const apiUri = "/api/v1"

func (a *Application) RegisterRouter() error {
	e := echo.New()

	group := e.Group(apiUri)
	a.checkHandler.Register(group)
	a.orderHandler.Register(group)

	a.router = e

	return nil
}
