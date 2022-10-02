package backup

import (
	"context"
	pb "intermediate_server/internal/models/pb"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/hashicorp/go.net/context"
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

	dbContext, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second*5))

	err = c.service.Create(dbContext, cancel, &b)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	err = c.service.SendBackupToMainServer(bodyBytes)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	backupID, err := c.service.GetLastID(dbContext)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	err = c.service.Delete(dbContext, backupID)

	return ctx.NoContent(http.StatusCreated)
}
