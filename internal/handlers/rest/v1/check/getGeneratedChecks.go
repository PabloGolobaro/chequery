package check

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type GeneratedChecksResponse struct {
	IDs []int `json:"ids"`
}

func (p *checkHandler) GetGeneratedChecks(ctx echo.Context) error {
	generatedChecksResponse, err := p.useCases.GetGeneratedCheckIDs(ctx.Request().Context())
	if err != nil {
		return echo.ErrInternalServerError
	}

	return ctx.JSON(http.StatusOK, generatedChecksResponse)
}
