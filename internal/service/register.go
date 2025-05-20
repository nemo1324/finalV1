package service

import (
	"context"
	"final/internal/repository/postgres/sqlc"
	finalv1 "final/pkg/proto/sync/final-boss/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strconv" // 👈 добавь
)

func (s *service) Register(ctx context.Context, req *finalv1.RegisterRequest) (*finalv1.RegisterResponse, error) {
	if err := req.ValidateAll(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	userID, err := s.DB.CreateUser(ctx, sqlc.CreateUserParams{
		Name:   req.GetUsername(),
		Login:  req.GetUsername(),
		Pass:   req.GetPassword(),
		Status: "register",
	})
	if err != nil {
		s.logger.Error("register error: %v", err)
		return nil, status.Error(codes.Internal, "internal error")
	}

	s.logger.Debug("created user id", "userID", userID)

	return &finalv1.RegisterResponse{
		UserId: strconv.FormatInt(int64(userID), 10), // 👈 теперь string
	}, nil
}
