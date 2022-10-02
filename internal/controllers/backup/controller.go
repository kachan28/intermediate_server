package backup

import (
	backupService "intermediate_server/internal/services/backup"

	validator "github.com/go-playground/validator/v10"
)

type Controller struct {
	validate *validator.Validate
	service  *backupService.Service
}

func InitController(dsn, mainServerURL string) (*Controller, error) {
	validate := validator.New()

	service, err := backupService.InitializeService(dsn, mainServerURL)
	if err != nil {
		return nil, err
	}

	return &Controller{validate: validate, service: service}, nil
}
