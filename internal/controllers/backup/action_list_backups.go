package backup

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (c *Controller) ListBackups(ctx echo.Context) error {
	folderName := ctx.QueryParam("folder")
	if folderName == "" {
		return ctx.String(http.StatusBadRequest, "folder param was not provided")
	}

	statusCode, body, err := c.service.ListBackups(folderName)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	MIMEType := echo.MIMETextPlainCharsetUTF8
	if statusCode == http.StatusOK {
		MIMEType = echo.MIMEApplicationJSON
	}

	return ctx.Blob(statusCode, MIMEType, body)
}
