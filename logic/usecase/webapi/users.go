package usecase

import (
	"context"
	"fmt"

	"github.com/stepundel1/E-commerce/Users/logic/entity"
	"github.com/stepundel1/E-commerce/pkg/auth"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCase struct {
	repo UserRepo
	jwt  *auth.JWTManager
}

func NewUserUseCase(repo UserRepo, jwt *auth.JWTManager) *UserUseCase {
	return &UserUseCase{
		repo: repo,
		jwt:  jwt,
	}
}

func (uc *UserUseCase) Register(ctx context.Context, email, name, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("UserUseCase - Register - bcrypt.GenerateFromPassword: %w", err)
	}

	user := entity.User{
		Email:        email,
		Name:         name,
		PasswordHash: string(hashedPassword),
	}

	err = uc.repo.Create(ctx, user)
	if err != nil {
		return fmt.Errorf("UserUseCase - Register - uc.repo.Create: %w", err)
	}

	return nil
}

func (uc *UserUseCase) Login(ctx context.Context, email, password string) (string, error) {
	user, err := uc.repo.GetByEmail(ctx, email)
	if err != nil {
		return "", fmt.Errorf("UserUseCase - Login - uc.repo.GetByEmail: %w", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return "", fmt.Errorf("UserUseCase - Login - bcrypt.CompareHashAndPassword: %w", err)
	}

	token, err := uc.jwt.Generate(user.ID)
	if err != nil {
		return "", fmt.Errorf("UserUseCase - Login - uc.jwt.Generate: %w", err)
	}

	return token, nil
}
