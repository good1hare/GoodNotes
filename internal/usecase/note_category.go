package usecase

import "MateMind/internal/entity"

type NoteCategoryUseCase struct {
	repo NoteCategoryRepo
}

func NewNoteCategoryUseCase(r NoteCategoryRepo) *NoteCategoryUseCase {
	return &NoteCategoryUseCase{
		repo: r,
	}
}

func (ncuc *NoteCategoryUseCase) GetNoteCategory(id int) (entity.NoteCategory, error) {
	note, err := ncuc.repo.FindNoteCategory(id)
	if err != nil {
		return entity.NoteCategory{}, err
	}
	return note, nil
}

func (ncuc *NoteCategoryUseCase) CreateNoteCategory(note entity.NoteCategory) (entity.NoteCategory, error) {
	note, err := ncuc.repo.SaveNoteCategory(note)
	if err != nil {
		return note, err
	}
	return note, nil
}

func (ncuc *NoteCategoryUseCase) UpdateNoteCategory(note entity.NoteCategory) (entity.NoteCategory, error) {
	note, err := ncuc.repo.UpdateNoteCategory(note)
	if err != nil {
		return entity.NoteCategory{}, err
	}
	return note, nil
}

func (ncuc *NoteCategoryUseCase) DeleteNoteCategory(id int) error {
	err := ncuc.repo.DeleteNoteCategory(id)
	if err != nil {
		return err
	}
	return nil
}
