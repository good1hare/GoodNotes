package usecase

import (
	"MateMind/internal/entity"
)

type UserUseCase struct {
	repo UserRepo
}

func NewUserUseCase(r UserRepo) *UserUseCase {
	return &UserUseCase{
		repo: r,
	}
}

func (u *UserUseCase) GetUser(chatID int) (entity.User, error) {
	user, err := u.repo.FindUser(chatID)
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (u *UserUseCase) CreateUser(user entity.User) (entity.User, error) {
	user, err := u.repo.SaveUser(user)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (u *UserUseCase) UpdateUser(user entity.User) (entity.User, error) {
	user, err := u.repo.UpdateUser(user)
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (u *UserUseCase) DeleteUser(chatID int) error {
	err := u.repo.DeleteUser(chatID)
	if err != nil {
		return err
	}
	return nil
}
