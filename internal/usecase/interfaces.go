package usecase

import (
	"MateMind/internal/entity"
)

type (
	User interface {
		GetUser(int) (entity.User, error)
		CreateUser(entity.User) (entity.User, error)
		UpdateUser(entity.User) (entity.User, error)
		DeleteUser(int) error
	}

	UserRepo interface {
		FindUser(int) (entity.User, error)
		SaveUser(entity.User) (entity.User, error)
		UpdateUser(entity.User) (entity.User, error)
		DeleteUser(int) error
	}

	Note interface {
		GetNote(int) (entity.Note, error)
		CreateNote(note entity.Note) (entity.Note, error)
		UpdateNote(note entity.Note) (entity.Note, error)
		DeleteNote(int) error
	}

	NoteRepo interface {
		FindNote(int) (entity.Note, error)
		SaveNote(note entity.Note) (entity.Note, error)
		UpdateNote(entity.Note) (entity.Note, error)
		DeleteNote(int) error
	}

	NoteCategory interface {
		GetNoteCategory(int) (entity.NoteCategory, error)
		CreateNoteCategory(entity.NoteCategory) (entity.NoteCategory, error)
		UpdateNoteCategory(entity.NoteCategory) (entity.NoteCategory, error)
		DeleteNoteCategory(int) error
	}

	NoteCategoryRepo interface {
		FindNoteCategory(int) (entity.NoteCategory, error)
		SaveNoteCategory(note entity.NoteCategory) (entity.NoteCategory, error)
		UpdateNoteCategory(entity.NoteCategory) (entity.NoteCategory, error)
		DeleteNoteCategory(int) error
	}
)
