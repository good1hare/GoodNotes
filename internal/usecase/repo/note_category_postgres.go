package repo

import (
	"MateMind/internal/entity"
	"MateMind/pkg/postgres"
	"context"
	"errors"
	"github.com/jackc/pgx/v4"
)

type NoteCategoryRepo struct {
	*postgres.Postgres
}

func NewNoteCategoryRepo(pg *postgres.Postgres) *NoteCategoryRepo {
	return &NoteCategoryRepo{pg}
}

func (r *NoteRepo) FindNoteCategory(id int) (entity.NoteCategory, error) {
	var category entity.NoteCategory
	row := r.Pool.QueryRow(context.TODO(), "SELECT * FROM note_categories WHERE id = $1", id)
	err := row.Scan(&category.Id, &category.ChatId, &category.Name, &category.ParentId, &category.Status)
	if err == pgx.ErrNoRows {
		return entity.NoteCategory{}, errors.New("entity not found")
	} else if err != nil {
		return entity.NoteCategory{}, err
	}
	return category, nil
}

func (r *NoteRepo) SaveNoteCategory(category entity.NoteCategory) (entity.NoteCategory, error) {
	_, err := r.Pool.Exec(context.TODO(),
		"INSERT INTO note_categories (chat_id, name, description, parent_id, status) VALUES ($1, $2, $3, $4)",
		category.ChatId, category.Name, category.ParentId, category.Status)
	if err != nil {
		return entity.NoteCategory{}, err
	}
	return entity.NoteCategory{}, nil
}

func (r *NoteRepo) UpdateNoteCategory(note entity.NoteCategory) (entity.NoteCategory, error) {
	_, err := r.Pool.Exec(
		context.TODO(),
		"UPDATE note_categories SET (chat_id=$1, name=$2, parent_id=$3, category_id=$4, status=$5) WHERE chat_id=$6",
		note.ChatId, note.Name, note.ParentId, note.Status,
	)
	if err != nil {
		return entity.NoteCategory{}, err
	}
	return entity.NoteCategory{}, nil
}

func (r *NoteRepo) DeleteNoteCategory(id int) error {
	_, err := r.Pool.Exec(context.TODO(), "DELETE FROM note_categories WHERE id=$1", id)
	return err
}
