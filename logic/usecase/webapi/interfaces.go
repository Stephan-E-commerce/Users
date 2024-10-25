package webapi

import (
	"context"

	"github.com/stepundel1/E-commerce/Users/logic/entity"
)

type (
	UserRepoInterface interface {
		Create(context.Context, entity.User) error
	}
)
