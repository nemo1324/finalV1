package final

import (
	"context"
	pb "final/pkg/proto/sync/final-boss/v1"
)

func (s *Server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	resp, err := s.svc.GetUser(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
