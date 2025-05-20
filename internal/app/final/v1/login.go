package final

import (
	"context"
	pb "final/pkg/proto/sync/final-boss/v1"
)

func (s *Server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	resp, err := s.svc.Login(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
