package usecase

import (
	"MateMind/internal/entity"
	"context"
)

// UserUseCase -.
type UserUseCase struct {
	repo UserRepo
}

func (o UserUseCase) GetUser(ctx context.Context, chatId int) (entity.User, error) {
	order, err := o.repo.FindUser(ctx, chatId)
	if err != nil {
		return entity.User{}, err
	}
	return order, nil
}

func (o UserUseCase) CreateUser(ctx context.Context) (entity.User, error) {
	e := entity.User{}

	return e, nil
}

func (o UserUseCase) DeleteUser(ctx context.Context, chatId int) error {
	return nil
}

func (o UserUseCase) UpdateUser(ctx context.Context) (entity.User, error) {
	e := entity.User{}

	return e, nil
}

// New -.
func New(r UserRepo) *UserUseCase {
	return &UserUseCase{
		repo: r,
	}
}
