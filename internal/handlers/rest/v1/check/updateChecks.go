package check

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func (c *checkHandler) UpdateChecksStatus(ctx echo.Context) error {
	queryIds := ctx.QueryParams()["id"]

	ids := make([]int, len(queryIds))

	for i, id := range queryIds {
		atoi, err := strconv.Atoi(id)
		if err != nil {
			return err
		}
		ids[i] = atoi
	}

	err := c.useCases.SetChecksStatusPrinted(ctx.Request().Context(), ids)
	if err != nil {
		return echo.ErrInternalServerError
	}

	return ctx.NoContent(http.StatusOK)
}
