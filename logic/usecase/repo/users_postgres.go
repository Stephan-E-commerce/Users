package repo

import (
	"context"
	"fmt"

	"github.com/stepundel1/E-commerce/Users/logic/entity"
	"github.com/stepundel1/E-commerce/pkg/postgres"
)

// UserRepo struct represents the repository for user-related database operations.
type UserRepo struct {
	*postgres.Postgres
}

// NewUserRepo creates a new instance of UserRepo.
func NewUserRepo(p *postgres.Postgres) *UserRepo {
	return &UserRepo{Postgres: p}
}

// Create inserts a new user into the database.
func (r *UserRepo) Create(ctx context.Context, user entity.User) error {

	// Build the SQL query to insert the user into the userlist table.
	sql, args, err := r.Builder.
		Insert("userlist").
		Columns("name", "email", "password_hash").
		Values(user.Name, user.Email, user.PasswordHash).
		ToSql()
	if err != nil {
		return fmt.Errorf("UserRepo - Create - r.Builder: %w", err)
	}

	// Execute the SQL query.
	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("UserRepo - Create - r.Pool.Exec: %w", err)
	}

	return nil
}
