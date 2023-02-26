package printer

import (
	"net/http"
)

const (
	url = "/api/v1/printer/:printer_id"
)

type UseCases interface {
	//GetGeneratedCheckIDs(ctx context.Context, printerID string) ([]string, error)
}

type printerHandler struct {
	useCases UseCases
}

func NewPrinterHandler(useCases UseCases) *printerHandler {
	return &printerHandler{useCases: useCases}
}

func (p *printerHandler) Register(mux http.ServeMux) {
	//mux.HandleFunc(url, p.GetGeneratedCheckIDs)
}
