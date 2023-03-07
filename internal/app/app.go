package app

import (
	"github.com/labstack/echo/v4"
	"github.com/pablogolobaro/chequery/internal/handlers"
	"go.uber.org/zap"
)

type Application struct {
	log            *zap.SugaredLogger
	router         *echo.Echo
	checkHandler   handlers.Handler
	orderHandler   handlers.Handler
	printerHandler handlers.Handler
}

func NewApplication(log *zap.SugaredLogger) *Application {
	return &Application{log: log}
}
