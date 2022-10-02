package backup

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (c *Controller) ListFolders(ctx echo.Context) error {
	folders, err := c.service.ListFolders()
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	return ctx.Blob(http.StatusOK, echo.MIMEApplicationJSON, folders)
}
