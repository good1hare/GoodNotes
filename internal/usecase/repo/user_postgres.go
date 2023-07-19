package repo

import (
	"MateMind/internal/entity"
	"MateMind/pkg/postgres"
	"context"
	"errors"
	"github.com/jackc/pgx/v4"
)

type UserRepo struct {
	*postgres.Postgres
}

func New(pg *postgres.Postgres) *UserRepo {
	return &UserRepo{pg}
}

func (r *UserRepo) FindUser(id int) (entity.User, error) {
	var user entity.User
	row := r.Pool.QueryRow(context.TODO(), "SELECT * FROM users WHERE id = $1", id)
	err := row.Scan(&user.Id, &user.UserName, &user.ChatId)
	if err == pgx.ErrNoRows {
		return entity.User{}, errors.New("entity not found")
	} else if err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (r *UserRepo) SaveUser(user entity.User) (entity.User, error) {
	_, err := r.Pool.Exec(context.TODO(),
		"INSERT INTO users (user_name, chat_id) VALUES ($1, $2) ON CONFLICT (chat_id) DO UPDATE SET user_name = excluded.user_name",
		user.UserName, user.ChatId)
	if err != nil {
		return entity.User{}, err
	}
	return entity.User{}, nil
}

func (r *UserRepo) UpdateUser(user entity.User) (entity.User, error) {
	_, err := r.Pool.Exec(context.TODO(), "UPDATE users SET user_name=$1 WHERE chat_id=$2", user.UserName, user.ChatId)
	if err != nil {
		return entity.User{}, err
	}
	return entity.User{}, nil
}

func (r *UserRepo) DeleteUser(chatID int) error {
	_, err := r.Pool.Exec(context.TODO(), "DELETE FROM users WHERE chat_id=$1", chatID)
	return err
}

//	r.db.Exec(ctx, "UPDATE user SET user_name = $1, chat_id = $2 WHERE id = $3",

//	r.db.Exec(ctx, "DELETE FROM user WHERE id = $1", id)
