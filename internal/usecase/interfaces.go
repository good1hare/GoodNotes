package usecase

import (
	"MateMind/internal/entity"
)

type (
	// User -
	User interface {
		GetUser(int) (entity.User, error)
		CreateUser(entity.User) (entity.User, error)
		UpdateUser(entity.User) (entity.User, error)
		DeleteUser(int) error
	}

	// UserRepo -
	UserRepo interface {
		FindUser(int) (entity.User, error)
		SaveUser(entity.User) (entity.User, error)
	}
)
