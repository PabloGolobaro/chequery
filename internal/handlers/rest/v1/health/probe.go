package health

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *healthCheckHandler) LiveProbe(ctx echo.Context) error {

	return ctx.JSON(http.StatusOK, LiveProbeResponse{Message: "I'm alive"})
}

type LiveProbeResponse struct {
	// The live message
	//
	// Required: true
	// Example: I'm alive
	Message string `json:"message"`
}
