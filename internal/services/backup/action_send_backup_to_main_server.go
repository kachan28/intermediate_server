package backup

import "fmt"

const saveApi = "/api/backup/save"

func (s *Service) SendBackupToMainServer(body []byte) error {
	resp, err := s.client.
		R().
		SetBody(body).
		Post(saveApi)

	if err != nil {
		return err
	}

	if !s.respIsSuccess(resp.StatusCode()) {
		return fmt.Errorf("wrong status code, err - %s", resp.Body())
	}

	return nil
}
