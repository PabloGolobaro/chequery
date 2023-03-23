package web

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"net/http"
)

const (
	urlIndex = "/"
)

type uiHandler struct {
	log *zap.SugaredLogger
}

func NewUiHandler(log *zap.SugaredLogger) *uiHandler {
	return &uiHandler{log: log}
}

func (u *uiHandler) Register(router *echo.Group) {
	router.Add(http.MethodGet, urlIndex, u.GetIndexPage)
}

func (u *uiHandler) GetIndexPage(c echo.Context) error {
	return c.Render(http.StatusOK, "base", nil)
}
