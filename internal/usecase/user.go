package usecase

import (
	"MateMind/internal/entity"
	"context"
)

// UserUseCase -.
type UserUseCase struct {
	repo UserRepo
}

func (o UserUseCase) GetOrder(ctx context.Context, chatId string) (entity.User, error) {
	order, err := o.repo.FindUser(ctx, chatId)
	if err != nil {
		return entity.User{}, err
	}
	return order, nil
}

func (o UserUseCase) CreateOrder(ctx context.Context) (entity.User, error) {
	e := entity.User{}

	return e, nil
}

// New -.
func New(r UserRepo) *UserUseCase {
	return &UserUseCase{
		repo: r,
	}
}
