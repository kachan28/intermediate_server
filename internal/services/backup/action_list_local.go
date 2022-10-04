package backup

import "intermediate_server/internal/models"

func (s *Service) ListLocal() (map[models.SystemID][]models.Backup, error) {
	return s.repository.List()
}
