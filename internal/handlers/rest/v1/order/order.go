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
	url = "/order"
)

type UseCases interface {
	CreateChecks(ctx context.Context, order entity.OrderDetails) error
}

type orderHandler struct {
	log           *zap.SugaredLogger
	checkUseCases UseCases
}

func NewOrderHandler(log *zap.SugaredLogger, checkUseCases UseCases) handlers.Handler {
	return &orderHandler{log: log, checkUseCases: checkUseCases}
}

func (o *orderHandler) Register(router *echo.Group) {
	router.Add(http.MethodPost, url, o.PostOrder)
}
