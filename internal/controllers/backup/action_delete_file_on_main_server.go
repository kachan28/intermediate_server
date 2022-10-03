package backup

import (
	"io/ioutil"
	"net/http"

	jsoniter "github.com/json-iterator/go"
	"github.com/labstack/echo/v4"
)

func (c *Controller) DeleteFileOnMainServer(ctx echo.Context) error {
	bodyBytes, err := ioutil.ReadAll(ctx.Request().Body)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	file := struct {
		FolderName string `json:"folder_name" validate:"required,gt=0"`
		FileName   string `json:"file_name" validate:"required,gt=0"`
	}{}
	err = jsoniter.Unmarshal(bodyBytes, &file)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	err = c.validate.Struct(&file)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	statusCode, body, err := c.service.DeleteFileOnMainServer(bodyBytes)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	MIMEType := echo.MIMETextPlainCharsetUTF8
	if statusCode == http.StatusOK {
		MIMEType = echo.MIMEApplicationJSON
	}

	return ctx.Blob(statusCode, MIMEType, body)
}
