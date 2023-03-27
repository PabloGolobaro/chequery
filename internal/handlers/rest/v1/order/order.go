package order

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/pablogolobaro/chequery/internal/domain/entity"
	"github.com/pablogolobaro/chequery/internal/handlers"
	"go.uber.org/zap"
	"net/http"
)

const (
	urlOrders = "/order"
)

type UseCases interface {
	CreateChecks(ctx context.Context, order entity.Order) (ids []int, err error)
}

type orderHandler struct {
	log      *zap.SugaredLogger
	useCases UseCases
}

func NewOrderHandler(log *zap.SugaredLogger, checkUseCases UseCases) handlers.Handler {
	return &orderHandler{log: log, useCases: checkUseCases}
}

func (o *orderHandler) Register(router *echo.Group) {
	router.Add(http.MethodPost, urlOrders, o.PostOrder)
}
