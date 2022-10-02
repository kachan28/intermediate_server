package backup

import (
	repo "intermediate_server/internal/repository"

	"github.com/go-resty/resty/v2"
)

type Service struct {
	client     *resty.Client
	repository *repo.Repository
}

func InitializeService(dsn, mainServerURL string) (*Service, error) {
	c := resty.New().SetBaseURL(mainServerURL)

	r, err := repo.InitRepo(dsn)
	if err != nil {
		return nil, err
	}
	return &Service{client: c, repository: r}, nil
}

func (s *Service) respIsSuccess(statusCode int) bool {
	return statusCode >= 200 && statusCode <= 299
}
