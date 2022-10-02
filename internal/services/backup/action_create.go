package backup

import (
	"context"
	pb "intermediate_server/internal/models/pb"
)

func (s *Service) Create(ctx context.Context, cancel context.CancelFunc, backup *pb.BackupCreate) error {
	return s.repository.Create(ctx, backup)
}
