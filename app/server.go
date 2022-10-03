package app

import (
	"fmt"
	"intermediate_server/config"
	backupController "intermediate_server/internal/controllers/backup"

	echoSwagger "github.com/swaggo/echo-swagger"

	_ "intermediate_server/docs"

	"github.com/labstack/echo/v4"
)

type Server struct {
	e *echo.Echo
}

func Init(conf *config.Config) (*Server, error) {
	e := echo.New()
	s := &Server{e: e}
	err := s.registerRoutes(conf.Database.DSN, conf.HTTP.MainServerURL)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (s *Server) Start(conf *config.Config) error {
	return s.e.Start(fmt.Sprintf(":%s", conf.HTTP.Port))
}

func (s *Server) registerRoutes(dsn, mainServerURL string) error {
	backupCtrl, err := backupController.InitController(dsn, mainServerURL)
	if err != nil {
		return err
	}

	//swagger doc handler
	s.e.GET("/swagger/*", echoSwagger.WrapHandler)

	s.e.POST("/api/backup/save", backupCtrl.Create)
	s.e.GET("/api/backup/list/folders", backupCtrl.ListFolders)
	s.e.GET("/api/backup/list/files", backupCtrl.ListBackups)
	s.e.POST("/api/backup/delete/remote/folder", backupCtrl.DeleteFolderOnMainServer)
	s.e.POST("/api/backup/delete/remote/file", backupCtrl.DeleteFileOnMainServer)
	return nil
}
