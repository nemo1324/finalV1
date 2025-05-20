package service

import (
	"context"
	finalv1 "final/pkg/proto/sync/final-boss/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *service) Logout(ctx context.Context, req *finalv1.LogoutRequest) (*finalv1.LogoutResponse, error) {
	if err := req.ValidateAll(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if err := s.DB.Logout(ctx, int32(req.GetUserId())); err != nil {
		return nil, status.Error(codes.Internal, "logout failed")
	}

	return &finalv1.LogoutResponse{}, nil
}
