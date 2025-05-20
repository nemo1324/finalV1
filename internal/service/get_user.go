package service

import (
	"context"
	finalv1 "final/pkg/proto/sync/final-boss/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *service) GetUser(ctx context.Context, req *finalv1.GetUserRequest) (*finalv1.GetUserResponse, error) {
	if err := req.ValidateAll(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	user, err := s.DB.GetUserByID(ctx, req.GetId())
	if err != nil {
		s.logger.Error("user not found: %v", err)
		return nil, status.Error(codes.NotFound, "user not found")
	}

	return &finalv1.GetUserResponse{
		Id:       int64(user.ID),
		Username: user.Name, // <-- здесь нужно достать имя из структуры
	}, nil
}
