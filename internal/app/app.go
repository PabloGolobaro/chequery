package app

import (
	"github.com/labstack/echo/v4"
	"github.com/pablogolobaro/chequery/internal/handlers"
	"github.com/pablogolobaro/chequery/pkg/renderer"
	"go.uber.org/zap"
)

type Application struct {
	log            *zap.SugaredLogger
	router         *echo.Echo
	checkHandler   handlers.Handler
	orderHandler   handlers.Handler
	printerHandler handlers.Handler
	healthHandler  handlers.Handler
	uiHandler      handlers.Handler
	authHandler    handlers.Handler
	renderer       *renderer.Template
}

func NewApplication(log *zap.SugaredLogger) *Application {
	return &Application{log: log}
}
