package repo

import (
	"context"
	"fmt"

	"github.com/stepundel1/E-commerce/Users/logic/entity"
	"github.com/stepundel1/E-commerce/pkg/postgres"
)

type UserRepo struct {
	*postgres.Postgres
}

func NewUserRepo(p *postgres.Postgres) *UserRepo {
	return &UserRepo{p}
}

func (r *UserRepo) Create(ctx context.Context, user entity.User) error {
	sql, args, err := r.Builder.
		Insert("users").
		Columns("email", "name", "password_hash").
		Values(user.Email, user.Name, user.PasswordHash).
		ToSql()
	if err != nil {
		return fmt.Errorf("UserRepo - Create - r.Builder: %w", err)
	}

	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("UserRepo - Create - r.Pool.Exec: %w", err)
	}

	return nil
}

func (r *UserRepo) GetByEmail(ctx context.Context, email string) (entity.User, error) {
	sql, args, err := r.Builder.
		Select("id", "email", "name", "password_hash").
		From("users").
		Where("email = ?", email).
		ToSql()
	if err != nil {
		return entity.User{}, fmt.Errorf("UserRepo - GetByEmail - r.Builder: %w", err)
	}

	var user entity.User
	err = r.Pool.QueryRow(ctx, sql, args...).Scan(&user.ID, &user.Email, &user.Name, &user.PasswordHash)
	if err != nil {
		return entity.User{}, fmt.Errorf("UserRepo - GetByEmail - r.Pool.QueryRow: %w", err)
	}

	return user, nil
}
