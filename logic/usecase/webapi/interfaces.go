package usecase

import (
	"context"

	"github.com/stepundel1/E-commerce/Users/logic/entity"
)

type (
	// Translation -.

	UserRepo interface {
		Create(context.Context, entity.User) error
		GetByEmail(context.Context, string) (entity.User, error)
	}
)
