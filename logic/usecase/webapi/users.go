package webapi

import (
	"context"
	"fmt"
	"log"
	"time"

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

	start := time.Now()
	err = uc.repo.Create(ctx, user)
	log.Printf("Database query execution time: %v", time.Since(start))
	if err != nil {
		return fmt.Errorf("UserUseCase - Register - uc.repo.Create: %w", err)
	}

	return nil
}

func (uc *UserUseCase) LogIn(ctx context.Context, email string, password string) error {

	user, err := uc.repo.GetByEmail(ctx, email)
	if err != nil {
		return fmt.Errorf("UserUseCase - LogIn - uc.repo.LogInSql: %w", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return fmt.Errorf("UserUseCase - LogIn - incorrect password: %w", err)
	}

	return nil
}
