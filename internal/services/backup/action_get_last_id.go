package backup

func (s *Service) GetLastID() (int64, error) {
	return s.repository.GetLastInsertedRowID()
}
