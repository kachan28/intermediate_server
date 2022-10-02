package backup

const (
	apiListBackups = "/api/backup/list/files"
	folderParam    = "folder"
)

func (s *Service) ListBackups(folderName string) (int, []byte, error) {
	resp, err := s.client.R().
		SetQueryParam(folderParam, folderName).
		Get(apiListBackups)

	if err != nil {
		return 0, nil, err
	}

	return resp.StatusCode(), resp.Body(), nil
}
