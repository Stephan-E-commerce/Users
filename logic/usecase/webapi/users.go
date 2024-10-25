package webapi

import (
	"context"
	"fmt"

	"github.com/stepundel1/E-commerce/Users/logic/entity"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCase struct {
	repo UserRepoInterface
}

func NewUserUseCase(repo UserRepoInterface) *UserUseCase {
	return &UserUseCase{
		repo: repo,
	}
}

// password hash
func (uc *UserUseCase) Register(ctx context.Context, user entity.User, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("UserUseCase - Register - bcrypt.GenerateFromPassword: %w", err)
	}

	user.PasswordHash = string(hashedPassword)

	err = uc.repo.Create(ctx, user)
	if err != nil {
		return fmt.Errorf("UserUseCase - Register - uc.repo.Create: %w", err)
	}

	return nil
}
