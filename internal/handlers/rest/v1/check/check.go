package check

import (
	"context"
	"net/http"
)

const (
	url = "/api/v1/check"
)

type UseCases interface {
	SetChecksStatusPrinted(ctx context.Context, checkID []string) error
}

type checkHandler struct {
	useCases UseCases
}

func NewCheckHandler(useCases UseCases) *checkHandler {
	return &checkHandler{useCases: useCases}
}

func (c *checkHandler) Register(mux http.ServeMux) {
	mux.HandleFunc(url, c.UpdateChecksStatus)
}
