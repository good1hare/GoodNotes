package repo

import (
	"MateMind/internal/entity"
	"MateMind/pkg/postgres"
	"context"
	"errors"
	"github.com/jackc/pgx/v4"
)

type NoteRepo struct {
	*postgres.Postgres
}

func NewNoteRepo(pg *postgres.Postgres) *NoteRepo {
	return &NoteRepo{pg}
}

func (r *NoteRepo) FindNote(id int) (entity.Note, error) {
	var note entity.Note
	row := r.Pool.QueryRow(context.TODO(), "SELECT * FROM notes WHERE id = $1", id)
	err := row.Scan(&note.Id, &note.ChatId, &note.Name, &note.Description, &note.CategoryId, &note.Status)
	if err == pgx.ErrNoRows {
		return entity.Note{}, errors.New("entity not found")
	} else if err != nil {
		return entity.Note{}, err
	}
	return note, nil
}

func (r *NoteRepo) SaveNote(note entity.Note) (entity.Note, error) {
	_, err := r.Pool.Exec(context.TODO(),
		"INSERT INTO notes (chat_id, name, description, category_id, status) VALUES ($1, $2, $3, $4, $5)",
		note.ChatId, note.Name, note.Description, note.CategoryId, note.Status)
	if err != nil {
		return entity.Note{}, err
	}
	return entity.Note{}, nil
}

func (r *NoteRepo) UpdateNote(note entity.Note) (entity.Note, error) {
	_, err := r.Pool.Exec(
		context.TODO(),
		"UPDATE notes SET (chat_id=$1, name=$2, description=$3, category_id=$4, status=$5) WHERE chat_id=$6",
		note.ChatId, note.Name, note.Description, note.CategoryId, note.Status, note.Id,
	)
	if err != nil {
		return entity.Note{}, err
	}
	return entity.Note{}, nil
}

func (r *NoteRepo) DeleteNote(id int) error {
	_, err := r.Pool.Exec(context.TODO(), "DELETE FROM notes WHERE id=$1", id)
	return err
}
