package final

import (
	"context"
	pb "final/pkg/proto/sync/final-boss/v1"
)

func (s *Server) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	resp, err := s.svc.Register(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
