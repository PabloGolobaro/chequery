package order

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/pablogolobaro/chequery/internal/domain/entity"
	"net/http"
	"strconv"
)

type OrderCreateRequest struct {
	Order string `json:"order"`
}

type OrderCreateResponse struct {
	// Example: [1,2,3]
	Ids []int `json:"ids"`
}

func (o *orderHandler) PostOrder(ctx echo.Context) error {
	queryParamPoint := ctx.QueryParam("point_id")
	if queryParamPoint == "" {

		return echo.ErrBadRequest
	}

	pointId, err := strconv.Atoi(queryParamPoint)
	if err != nil {
		return echo.ErrBadRequest
	}

	request := OrderCreateRequest{}

	err = ctx.Bind(&request)
	if err != nil {

		return echo.ErrBadRequest
	}

	if !json.Valid([]byte(request.Order)) {
		return echo.ErrBadRequest
	}

	ids, err := o.useCases.CreateChecks(ctx.Request().Context(), entity.OrderDetails{PointID: pointId, Order: request.Order})
	if err != nil {

		return echo.ErrInternalServerError
	}

	return ctx.JSON(http.StatusCreated, OrderCreateResponse{Ids: ids})
}
