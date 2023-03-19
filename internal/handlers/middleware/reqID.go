package middleware

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func GetCurrencyReqID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		reqID := uuid.New().String()
		c.Set("req-id", reqID)
		return next(c)
	}
}
