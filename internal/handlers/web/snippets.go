package web

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (u *uiHandler) PostAddProductForm(c echo.Context) error {
	return c.Render(http.StatusOK, addProductSnippet, nil)
}
