package check

import (
	"context"
	"github.com/labstack/echo/v4"
	"net/http"
)

const (
	urlGetPDF       = "/check/:check_id/pdf"
	url             = "/check"
	urlGetGenerated = "/check/generated"
)

type UseCases interface {
	GetGeneratedCheckIDs(ctx context.Context) (IDs, error)
	SetChecksStatusPrinted(ctx context.Context, checkIDs []string) error
	GetCheckFilePath(ctx context.Context, checkID string) (string, error)
}

type checkHandler struct {
	useCases UseCases
}

func NewCheckHandler(useCases UseCases) *checkHandler {
	return &checkHandler{useCases: useCases}
}

func (c *checkHandler) Register(router echo.Router) {
	router.Add(http.MethodGet, urlGetPDF, c.GetCheckPDF)
	router.Add(http.MethodGet, urlGetGenerated, c.GetGeneratedChecks)
	router.Add(http.MethodPut, url, c.UpdateChecksStatus)
}
