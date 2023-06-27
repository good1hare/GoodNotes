package repo

import (
	"MateMind/internal/entity"
	"context"
	"errors"
	"github.com/jackc/pgx/v4"
)

type UserRepo struct {
	db *pgx.Conn
}

func New(db *pgx.Conn) *UserRepo {
	return &UserRepo{db}
}

func (r *UserRepo) FindOrder(ctx context.Context, id int) (entity.User, error) {
	var user entity.User
	row := r.db.QueryRow(ctx, "SELECT * FROM user WHERE id = $1", id)
	err := row.Scan(&user.Id, &user.UserName, &user.ChatId)
	if err == pgx.ErrNoRows {
		return entity.User{}, errors.New("entity not found")
	} else if err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (r *UserRepo) SaveOrder(ctx context.Context, user entity.User) error {
	_, err := r.db.Exec(ctx, "INSERT INTO user (user_name, chat_id) VALUES ($1, $2)",
		user.UserName, user.ChatId)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepo) UpdateOrder(ctx context.Context, id int, user entity.User) error {
	_, err := r.db.Exec(ctx, "UPDATE user SET user_name = $1, chat_id = $2 WHERE id = $3",
		user.UserName, user.ChatId, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepo) DeleteOrder(ctx context.Context, id int) error {
	_, err := r.db.Exec(ctx, "DELETE FROM user WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
