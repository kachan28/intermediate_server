package backup

const deleteFileApi = "/api/backup/delete/file"

func (s *Service) DeleteFileOnMainServer(body []byte) (int, []byte, error) {
	resp, err := s.client.R().SetBody(body).Post(deleteFileApi)
	if err != nil {
		return 0, nil, err
	}

	return resp.StatusCode(), resp.Body(), nil
}
