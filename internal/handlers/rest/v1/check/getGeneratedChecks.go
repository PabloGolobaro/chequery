package check

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type IDs []string

func (p *checkHandler) GetGeneratedChecks(ctx echo.Context) error {
	checkIDs, err := p.useCases.GetGeneratedCheckIDs(ctx.Request().Context())
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, checkIDs)
}
