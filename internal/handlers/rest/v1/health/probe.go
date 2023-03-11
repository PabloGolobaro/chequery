package health

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *healthCheckHandler) LiveProbe(ctx echo.Context) error {

	return ctx.JSON(http.StatusOK, "I'm alive")
}
