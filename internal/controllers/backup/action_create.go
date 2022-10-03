package backup

import (
	pb "intermediate_server/internal/models/pb"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"
	"google.golang.org/protobuf/proto"
)

func (c *Controller) Create(ctx echo.Context) error {
	bodyBytes, err := ioutil.ReadAll(ctx.Request().Body)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	b := pb.BackupCreate{}
	err = proto.Unmarshal(bodyBytes, &b)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	err = c.validate.Struct(&b)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	err = c.service.Create(&b)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	err = c.service.SendBackupToMainServer(bodyBytes)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	backupID, err := c.service.GetLastID()
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	err = c.service.Delete(backupID)

	return ctx.NoContent(http.StatusCreated)
}
