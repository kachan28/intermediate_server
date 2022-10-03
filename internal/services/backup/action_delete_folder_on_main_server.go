package backup

const deleteFolderApi = "/api/backup/delete/folder"

func (s *Service) DeleteFolderOnMainServer(body []byte) (int, []byte, error) {
	resp, err := s.client.R().SetBody(body).Post(deleteFolderApi)
	if err != nil {
		return 0, nil, err
	}

	return resp.StatusCode(), resp.Body(), nil
}
