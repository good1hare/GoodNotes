package usecase

import (
	"MateMind/internal/entity"
)

type NoteUseCase struct {
	repo NoteRepo
}

func NewNoteUseCase(r NoteRepo) *NoteUseCase {
	return &NoteUseCase{
		repo: r,
	}
}

func (nuc *NoteUseCase) GetNote(id int) (entity.Note, error) {
	note, err := nuc.repo.FindNote(id)
	if err != nil {
		return entity.Note{}, err
	}
	return note, nil
}

func (nuc *NoteUseCase) CreateNote(note entity.Note) (entity.Note, error) {
	note, err := nuc.repo.SaveNote(note)
	if err != nil {
		return note, err
	}
	return note, nil
}

func (nuc *NoteUseCase) UpdateNote(note entity.Note) (entity.Note, error) {
	note, err := nuc.repo.UpdateNote(note)
	if err != nil {
		return entity.Note{}, err
	}
	return note, nil
}

func (nuc *NoteUseCase) DeleteNote(id int) error {
	err := nuc.repo.DeleteNote(id)
	if err != nil {
		return err
	}
	return nil
}
