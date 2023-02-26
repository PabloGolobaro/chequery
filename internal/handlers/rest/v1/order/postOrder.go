package order

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
)

func (o *orderHandler) PostOrder(ctx echo.Context) error {

	body := ctx.Request().Body
	defer body.Close()

	orderBytes, err := io.ReadAll(body)
	if err != nil {
		return err
	}

	if !json.Valid(orderBytes) {
		return err
	}

	err = o.checkUseCases.CreateChecks(ctx.Request().Context(), string(orderBytes))
	if err != nil {

		return err
	}

	return ctx.JSON(http.StatusCreated, "Order registered successfully")
}
