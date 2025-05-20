package final

import (
	"context"
	pb "final/pkg/proto/sync/final-boss/v1"
)

func (s *Server) Logout(ctx context.Context, req *pb.LogoutRequest) (*pb.LogoutResponse, error) {
	resp, err := s.svc.Logout(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
