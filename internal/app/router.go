package app

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pablogolobaro/chequery/internal/config"
	"github.com/pablogolobaro/chequery/internal/handlers/auth"
)

const (
	apiUri = "/api/v1"
)

func (a *Application) RegisterRouter(conf config.Config) error {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	c := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(auth.JwtCustomClaims)
		},
		SigningKey: []byte(conf.JWTSecret),
	}

	e.Static("/", "static")

	e.Renderer = a.renderer

	openGroup := e.Group("")

	a.authHandler.Register(openGroup)

	a.uiHandler.Register(openGroup)

	a.healthHandler.Register(openGroup)

	apiGroup := openGroup.Group(apiUri)

	apiGroup.Use(echojwt.WithConfig(c))

	a.orderHandler.Register(apiGroup)

	a.checkHandler.Register(apiGroup)

	a.router = e

	return nil
}
