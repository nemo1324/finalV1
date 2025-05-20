package final

import (
	"context"
	pb "final/pkg/proto/sync/final-boss/v1"
)

func (s *Server) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	resp, err := s.svc.DeleteUser(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
