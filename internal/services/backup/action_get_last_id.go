package backup

import "context"

func (s *Service) GetLastID(ctx context.Context) (int64, error) {
	return s.repository.GetLastInsertedRowID(ctx)
}
