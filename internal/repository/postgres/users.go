package postgres

import (
	"context"
	"final/internal/repository/postgres/sqlc"
)

func (p *Postgres) CreateUser(ctx context.Context, arg sqlc.CreateUserParams) (int32, error) {
	// Логируем, чтобы убедиться, что статус реально передаётся
	p.logger.Debug("CreateUser called", "login", arg.Login, "status", arg.Status)

	// sqlc ожидает *CreateUserParams — передаём по указателю
	return p.queries.CreateUser(ctx, &arg)
}

func (p *Postgres) GetUserByLogin(ctx context.Context, login string) (*sqlc.User, error) {
	return p.queries.GetUserByLogin(ctx, login)
}

func (p *Postgres) GetUserByID(ctx context.Context, id int64) (*sqlc.User, error) {
	return p.queries.GetUserByID(ctx, int32(id))
}

func (p *Postgres) UpdateUserStatus(ctx context.Context, arg sqlc.UpdateUserStatusParams) error {
	return p.queries.UpdateUserStatus(ctx, &arg)
}

func (p *Postgres) UpdateUserName(ctx context.Context, arg sqlc.UpdateUserNameParams) error {
	return p.queries.UpdateUserName(ctx, &arg)
}

func (p *Postgres) DeleteUser(ctx context.Context, id int64) error {
	return p.queries.DeleteUser(ctx, int32(id))
}

func (p *Postgres) Logout(ctx context.Context, id int32) error {
	return p.queries.Logout(ctx, id)
}
