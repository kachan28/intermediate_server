package backup

import (
	"fmt"
)

const listFoldersApi = "/api/backup/list/folders"

func (s *Service) ListFolders() ([]byte, error) {
	resp, err := s.client.R().Get(listFoldersApi)
	if err != nil {
		return nil, err
	}

	if !s.respIsSuccess(resp.StatusCode()) {
		return nil, fmt.Errorf("wrong status code, err - %s", resp.Body())
	}

	return resp.Body(), nil
}
