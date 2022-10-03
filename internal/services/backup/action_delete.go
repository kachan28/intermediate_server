package backup

func (s *Service) Delete(backupID int64) error {
	return s.repository.Delete(backupID)
}
