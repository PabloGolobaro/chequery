package check

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/pablogolobaro/chequery/internal/handlers"
	"go.uber.org/zap"
	"net/http"
)

const (
	urlGetPDF       = "/check/:check_id/pdf"
	urlCheck        = "/check"
	urlGetGenerated = "/check/generated"
)

type UseCases interface {
	GetGeneratedCheckIDs(ctx context.Context) (GeneratedChecksResponse, error)
	SetChecksStatusPrinted(ctx context.Context, checkIDs []int) error
	GetCheckFilePath(ctx context.Context, checkID int) (string, error)
}

type checkHandler struct {
	log      *zap.SugaredLogger
	useCases UseCases
}

func NewCheckHandler(log *zap.SugaredLogger, useCases UseCases) handlers.Handler {
	return &checkHandler{log: log, useCases: useCases}
}

func (c *checkHandler) Register(router *echo.Group) {
	router.Add(http.MethodGet, urlGetPDF, c.GetCheckPDF)
	router.Add(http.MethodGet, urlGetGenerated, c.GetGeneratedChecks)
	router.Add(http.MethodPut, urlCheck, c.UpdateChecksStatus)
}
