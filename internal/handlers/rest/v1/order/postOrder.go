package order

import (
	"encoding/json"
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/pablogolobaro/chequery/internal/domain/entity"
	"net/http"
	"strconv"
)

const (
	queryError = "wrong query parameter"
	bodyError  = "wrong body parameters"
	orderError = "order is not a valid json"
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

		return echo.ErrBadRequest.SetInternal(errors.New(queryError))
	}

	pointId, err := strconv.Atoi(queryParamPoint)
	if err != nil {
		return echo.ErrBadRequest.SetInternal(errors.New(queryError))
	}

	request := OrderCreateRequest{}

	err = ctx.Bind(&request)
	if err != nil {

		return echo.ErrBadRequest.SetInternal(errors.New(bodyError))
	}

	if !json.Valid([]byte(request.Order)) {
		return echo.ErrBadRequest.SetInternal(errors.New(orderError))
	}

	ids, err := o.useCases.CreateChecks(ctx.Request().Context(), entity.OrderDetails{PointID: pointId, Order: request.Order})
	if err != nil {

		return echo.ErrInternalServerError
	}

	return ctx.JSON(http.StatusCreated, OrderCreateResponse{Ids: ids})
}
