package final

import (
	"context"
	pb "final/pkg/proto/sync/final-boss/v1"
)

func (s *Server) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	resp, err := s.svc.UpdateUser(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
