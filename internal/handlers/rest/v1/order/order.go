package order

import (
	"context"
	"net/http"
)

const (
	url = "/api/v1/order"
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

func (o *orderHandler) Register(mux http.ServeMux) {
	mux.HandleFunc(url, o.PostOrder)
}
