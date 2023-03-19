package check

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"strings"
)

func (c *checkHandler) UpdateChecksStatus(ctx echo.Context) error {
	handledParams := make([]string, 0)

	queryIds := ctx.QueryParams()["id"]
	for _, queryParam := range queryIds {
		strings.Trim(queryParam, "\"")
		split := strings.Split(queryParam, ",")
		handledParams = append(handledParams, split...)
	}

	ids := make([]int, len(handledParams))

	for i, id := range handledParams {
		atoi, err := strconv.Atoi(id)
		if err != nil {
			return echo.ErrBadRequest.WithInternal(err)
		}
		ids[i] = atoi
	}

	err := c.useCases.SetChecksStatusPrinted(ctx.Request().Context(), ids)
	if err != nil {
		c.log.Errorw("Check.useCases.SetChecksStatusPrinted", "error: ", err, "ids", ids)
		return echo.ErrInternalServerError
	}

	return ctx.NoContent(http.StatusOK)
}
