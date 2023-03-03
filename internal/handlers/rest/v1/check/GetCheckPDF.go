package check

import (
	"github.com/labstack/echo/v4"
	"strconv"
)

func (c checkHandler) GetCheckPDF(ctx echo.Context) error {
	checkId := ctx.Param("check_id")

	id, err := strconv.Atoi(checkId)
	if err != nil {
		return err
	}

	filePath, err := c.useCases.GetCheckFilePath(ctx.Request().Context(), id)
	if err != nil {
		return err
	}

	return ctx.File(filePath)
}
