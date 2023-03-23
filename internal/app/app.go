package app

import (
	"github.com/labstack/echo/v4"
	"github.com/pablogolobaro/chequery/internal/handlers"
	"go.uber.org/zap"
	"html/template"
)

type Application struct {
	log            *zap.SugaredLogger
	router         *echo.Echo
	checkHandler   handlers.Handler
	orderHandler   handlers.Handler
	printerHandler handlers.Handler
	healthHandler  handlers.Handler
	uiHandler      handlers.Handler
	t              *template.Template
}

func NewApplication(log *zap.SugaredLogger) *Application {
	return &Application{log: log}
}
