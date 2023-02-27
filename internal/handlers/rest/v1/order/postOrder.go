package order

import (
	"github.com/labstack/echo/v4"
	"github.com/pablogolobaro/chequery/internal/domain/entity"
	"net/http"
)

func (o *orderHandler) PostOrder(ctx echo.Context) error {
	var order entity.OrderDetails

	err := ctx.Bind(&order)
	if err != nil {
		return err
	}

	err = o.checkUseCases.CreateChecks(ctx.Request().Context(), order)
	if err != nil {

		return err
	}

	return ctx.JSON(http.StatusCreated, "OrderDetails registered successfully")
}
