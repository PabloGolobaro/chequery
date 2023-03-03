package printer

import (
	"github.com/labstack/echo/v4"
	"github.com/pablogolobaro/chequery/internal/handlers"
	"go.uber.org/zap"
)

const (
	url = "/api/v1/printer/:printer_id"
)

type UseCases interface {
}

type printerHandler struct {
	log      zap.SugaredLogger
	useCases UseCases
}

func NewPrinterHandler(log zap.SugaredLogger, useCases UseCases) handlers.Handler {
	return &printerHandler{log: log, useCases: useCases}
}

func (p *printerHandler) Register(router *echo.Group) {
}
