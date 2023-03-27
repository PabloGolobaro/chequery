package order

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/pablogolobaro/chequery/internal/domain/entity"
	"net/http"
)

const (
	bodyError = "wrong body parameters"
)

type OrderCreateRequest struct {
	Order entity.Order `json:"order"`
}

type OrderCreateResponse struct {
	// Example: [1,2,3]
	Ids []int `json:"ids"`
}

func (o *orderHandler) PostOrder(ctx echo.Context) error {
	request := OrderCreateRequest{}

	err := ctx.Bind(&request.Order)
	if err != nil {

		return echo.ErrBadRequest.SetInternal(errors.New(bodyError))
	}

	ids, err := o.useCases.CreateChecks(ctx.Request().Context(), entity.Order{PointID: request.Order.PointID, Products: request.Order.Products})
	if err != nil {

		return echo.ErrInternalServerError
	}

	return ctx.JSON(http.StatusCreated, OrderCreateResponse{Ids: ids})
}
