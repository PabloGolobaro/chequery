package app

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pablogolobaro/chequery/internal/handlers/auth"
)

const (
	jwtSecret = "secret"
	apiUri    = "/api/v1"
)

func (a *Application) RegisterRouter() error {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(auth.JwtCustomClaims)
		},
		SigningKey: []byte(jwtSecret),
	}
	e.Use(echojwt.WithConfig(config))

	e.Static("/", "static")

	e.Renderer = a.renderer

	uiGroup := e.Group("")

	a.uiHandler.Register(uiGroup)

	apiGroup := uiGroup.Group(apiUri)

	a.checkHandler.Register(apiGroup)

	a.orderHandler.Register(apiGroup)

	a.healthHandler.Register(apiGroup)

	a.authHandler.Register(apiGroup)

	a.router = e

	return nil
}
