package service

import (
	"context"
	"final/internal/security/jwt"
	finalv1 "final/pkg/proto/sync/final-boss/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strconv"
)

func (s *service) Login(ctx context.Context, req *finalv1.LoginRequest) (*finalv1.LoginResponse, error) {
	if err := req.ValidateAll(); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	user, err := s.DB.GetUserByLogin(ctx, req.GetUsername())
	if err != nil {
		s.logger.Error("login failed: %v", err)
		return nil, status.Error(codes.NotFound, "invalid credentials")
	}

	if user.Pass != req.GetPassword() { // ⚠️ обязательно заменить на hash в проде!
		return nil, status.Error(codes.Unauthenticated, "invalid credentials")
	}

	// ✅ Генерация настоящего JWT токена
	token, err := jwt.GenerateAccessToken(strconv.Itoa(int(user.ID)), "user")
	if err != nil {
		s.logger.Error("failed to generate JWT token", "err", err)
		return nil, status.Error(codes.Internal, "token generation failed")
	}

	return &finalv1.LoginResponse{
		Token:  token,
		UserId: int64(user.ID),
	}, nil
}
