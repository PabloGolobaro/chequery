package web

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (u *uiHandler) GetHomePage(c echo.Context) error {
	return c.Render(http.StatusOK, "home", nil)
}
