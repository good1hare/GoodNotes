package usecase

import (
	"context"

	"MateMind/internal/entity"
)

type (
	// User -
	User interface {
		GetUser(context.Context, int) (entity.User, error)
		CreateUser(context.Context) (entity.User, error)
		UpdateUser(context.Context) (entity.User, error)
		DeleteUser(context.Context, int) error
	}

	// UserRepo -
	UserRepo interface {
		FindUser(context.Context, int) (entity.User, error)
		SaveUser(context.Context, entity.User) (entity.User, error)
	}
)
