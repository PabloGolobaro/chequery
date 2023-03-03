package order

import (
	"github.com/labstack/echo/v4"
	"github.com/pablogolobaro/chequery/internal/domain/entity"
	"net/http"
	"strconv"
)

func (o *orderHandler) PostOrder(ctx echo.Context) error {
	queryParamPoint := ctx.QueryParam("point_id")
	if queryParamPoint == "" {

		return echo.ErrBadRequest
	}
	pointId, err := strconv.Atoi(queryParamPoint)
	if err != nil {
		return echo.ErrBadRequest
	}

	var order = entity.OrderDetails{PointID: pointId, M: map[string]interface{}{}}

	err = ctx.Bind(&order.M)
	if err != nil {

		return echo.ErrBadRequest
	}

	err = o.checkUseCases.CreateChecks(ctx.Request().Context(), order)
	if err != nil {

		return echo.ErrInternalServerError
	}

	return ctx.JSON(http.StatusCreated, "OrderDetails registered successfully")
}
