package usecase

import (
	"context"

	"MateMind/internal/entity"
)

type (
	// User -
	User interface {
		GetUser(context.Context) (entity.User, error)
		CreateUser(context.Context) (entity.User, error)
		UpdateUser(context.Context) (entity.User, error)
		DeleteUser(context.Context) error
	}

	// UserRepo -
	UserRepo interface {
		FindUser(context.Context, string) (entity.User, error)
		SaveUser(context.Context) (entity.User, error)
	}
)
