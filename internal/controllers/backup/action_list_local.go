package backup

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (c *Controller) ListLocal(ctx echo.Context) error {
	backups, err := c.service.ListLocal()
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, backups)
}
