package backup

import (
	pb "intermediate_server/internal/models/pb"
)

func (s *Service) Create(backup *pb.BackupCreate) error {
	return s.repository.Create(backup)
}
