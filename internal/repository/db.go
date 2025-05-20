package repository

import (
	"context"
	"final/internal/repository/postgres/sqlc"
)

type DB interface {
	CreateUser(ctx context.Context, arg sqlc.CreateUserParams) (int32, error)
	GetUserByLogin(ctx context.Context, login string) (*sqlc.User, error)
	GetUserByID(ctx context.Context, id int64) (*sqlc.User, error)
	UpdateUserStatus(ctx context.Context, arg sqlc.UpdateUserStatusParams) error
	UpdateUserName(ctx context.Context, arg sqlc.UpdateUserNameParams) error
	DeleteUser(ctx context.Context, id int64) error
	Logout(ctx context.Context, id int32) error
}
