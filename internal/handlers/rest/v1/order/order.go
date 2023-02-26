package order

import (
	"context"
	"github.com/labstack/echo/v4"
	"net/http"
)

const (
	url = "/order"
)

type UseCases interface {
	CreateChecks(ctx context.Context, order string) error
}

type orderHandler struct {
	checkUseCases UseCases
}

func NewOrderHandler(checkUseCases UseCases) *orderHandler {
	return &orderHandler{checkUseCases: checkUseCases}
}

func (o *orderHandler) Register(router echo.Router) {
	router.Add(http.MethodPost, url, o.PostOrder)
}
