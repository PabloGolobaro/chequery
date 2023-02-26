package check

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (c *checkHandler) UpdateChecksStatus(ctx echo.Context) error {
	ids := ctx.QueryParams()["ids"]

	err := c.useCases.SetChecksStatusPrinted(ctx.Request().Context(), ids)
	if err != nil {
		return err
	}

	return ctx.NoContent(http.StatusOK)
}
