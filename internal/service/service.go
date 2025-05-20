package service

import (
	"context"
	"final/internal/repository"
	"final/internal/utils/observability/log"
	finalv1 "final/pkg/proto/sync/final-boss/v1"
)

type Service interface {
	Login(ctx context.Context, req *finalv1.LoginRequest) (*finalv1.LoginResponse, error)
	Register(ctx context.Context, req *finalv1.RegisterRequest) (*finalv1.RegisterResponse, error)
	Logout(ctx context.Context, req *finalv1.LogoutRequest) (*finalv1.LogoutResponse, error)
	GetUser(ctx context.Context, req *finalv1.GetUserRequest) (*finalv1.GetUserResponse, error)
	UpdateUser(ctx context.Context, req *finalv1.UpdateUserRequest) (*finalv1.UpdateUserResponse, error)
	DeleteUser(ctx context.Context, req *finalv1.DeleteUserRequest) (*finalv1.DeleteUserResponse, error)
}

type service struct {
	logger *log.Logger
	DB     repository.DB
}

func NewService(
	logger *log.Logger,
	DB repository.DB,
) Service {
	return &service{
		logger: logger,
		DB:     DB,
	}
}
