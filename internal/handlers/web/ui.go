package web

import (
	"github.com/labstack/echo/v4"
	"github.com/pablogolobaro/chequery/internal/handlers/rest/v1/order"
	"go.uber.org/zap"
	"net/http"
)

const (
	homepageUri       = "/"
	makeOrderURL      = "/order"
	addProductFormURL = "/add_product_form"
	addProductSnippet = "snippet_add_product"
)

type uiHandler struct {
	log      *zap.SugaredLogger
	useCases order.UseCases
}

func NewUiHandler(log *zap.SugaredLogger, useCases order.UseCases) *uiHandler {
	return &uiHandler{log: log, useCases: useCases}
}

func (u *uiHandler) Register(router *echo.Group) {
	router.Add(http.MethodGet, homepageUri, u.GetHomePage)
	router.Add(http.MethodGet, makeOrderURL, u.GetMakeOrder)
	router.Add(http.MethodPost, makeOrderURL, u.PostMakeOrder)
	router.Add(http.MethodPost, addProductFormURL, u.PostAddProductForm)
}
