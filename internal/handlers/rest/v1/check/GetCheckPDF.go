package check

import "github.com/labstack/echo/v4"

func (c checkHandler) GetCheckPDF(ctx echo.Context) error {
	checkId := ctx.Param("check_id")

	filePath, err := c.useCases.GetCheckFilePath(ctx.Request().Context(), checkId)
	if err != nil {
		return err
	}

	return ctx.File(filePath)
}
