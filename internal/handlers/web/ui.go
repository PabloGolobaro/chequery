package web

import (
	"github.com/labstack/echo/v4"
	"github.com/pablogolobaro/chequery/internal/domain/entity"
	"github.com/pablogolobaro/chequery/internal/handlers/rest/v1/order"
	"go.uber.org/zap"
	"io"
	"net/http"
	"os"
	"strconv"
)

const (
	makeOrderURL       = "/order"
	addProductFormURL  = "/add_product_form"
	addProductFormFile = "./static/templates/order_form/add_product.html"
)

type uiHandler struct {
	log      *zap.SugaredLogger
	useCases order.UseCases
}

func NewUiHandler(log *zap.SugaredLogger, useCases order.UseCases) *uiHandler {
	return &uiHandler{log: log, useCases: useCases}
}

func (u *uiHandler) Register(router *echo.Group) {
	router.Add(http.MethodGet, makeOrderURL, u.GetMakeOrder)
	router.Add(http.MethodPost, makeOrderURL, u.PostMakeOrder)
	router.Add(http.MethodPost, addProductFormURL, u.PostAddProductForm)
}

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

	order := entity.Order{
		PointID:  point_id,
		Products: products,
	}

	createChecks, err := u.useCases.CreateChecks(c.Request().Context(), order)
	if err != nil {
		return echo.ErrInternalServerError
	}

	var data = map[string]interface{}{"ids": createChecks}

	return c.Render(http.StatusOK, "order_created", data)
}

func (u *uiHandler) PostAddProductForm(c echo.Context) error {
	open, err := os.Open(addProductFormFile)
	if err != nil {
		return err
	}
	all, err := io.ReadAll(open)
	if err != nil {
		return err
	}
	return c.HTML(http.StatusOK, string(all))
}
