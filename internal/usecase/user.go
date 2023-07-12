package usecase

import (
	"MateMind/internal/entity"
)

// UserUseCase -.
type UserUseCase struct {
	repo UserRepo
}

// NewUserUseCase -.
func NewUserUseCase(r UserRepo) *UserUseCase {
	return &UserUseCase{
		repo: r,
	}
}

func (u UserUseCase) GetUser(chatId int) (entity.User, error) {
	user, err := u.repo.FindUser(chatId)
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (u UserUseCase) CreateUser(user entity.User) (entity.User, error) {

	return user, nil
}

func (u UserUseCase) UpdateUser(user entity.User) (entity.User, error) {

	return user, nil
}

func (u UserUseCase) DeleteUser(chatId int) error {
	return nil
}
