package web

import (
	"github.com/labstack/echo/v4"
	"github.com/pablogolobaro/chequery/internal/domain/entity"
	"net/http"
	"strconv"
)

func (u *uiHandler) GetMakeOrder(c echo.Context) error {
	return c.Render(http.StatusOK, "order_form", nil)
}

func (u *uiHandler) PostMakeOrder(c echo.Context) error {
	params, err := c.FormParams()
	if err != nil {
		return err
	}

	point_id, err := strconv.Atoi(params.Get("point_id"))
	if err != nil {
		return echo.ErrBadRequest
	}

	products := make([]entity.Product, len(params["name"]))

	for key, value := range params {
		switch key {
		case "name":
			for i, name := range value {
				products[i].Name = name
			}
		case "price":
			for i, price := range value {
				atoi, _ := strconv.Atoi(price)
				products[i].Price = atoi
			}
		case "quantity":
			for i, quantity := range value {
				atoi, _ := strconv.Atoi(quantity)
				products[i].Quantity = atoi
			}
		default:
			continue
		}
	}

	newOrder := entity.Order{
		PointID:  point_id,
		Products: products,
	}

	createChecks, err := u.useCases.CreateChecks(c.Request().Context(), newOrder)
	if err != nil {
		return echo.ErrInternalServerError
	}

	var data = map[string]interface{}{"ids": createChecks}

	return c.Render(http.StatusOK, "order_created", data)
}
