package health

import (
	"github.com/labstack/echo/v4"
	"github.com/pablogolobaro/chequery/internal/handlers"
	"net/http"
)

const (
	url = "/health-check"
)

type healthCheckHandler struct {
}

func NewHealthCheckHandler() handlers.Handler {
	return &healthCheckHandler{}
}

func (h *healthCheckHandler) Register(router *echo.Group) {
	router.Add(http.MethodGet, url, h.LiveProbe)
}
