package repositories

import (
  "fmt"
  e "github.com/robsonandradev/notes_api/entities"
)

type INoteRepository interface {
  CreateNote(author, title, content string) (n e.Note, err error)
  GetNoteByTitle(title string) (n []e.Note, err error)
  GetNotesByAuthor(author string) (n []e.Note, err error)
  GetNoteByAuthorAndTitle(author, title string) (n []e.Note, err error)
}

type NoteRepository struct {
  PG DBConnection
}

func NewNoteRepository(connector string) (*NoteRepository, error) {
  if connector == "postgres" {
    ur := &NoteRepository{ PG: &PostgresCon{} }
    return ur, nil
  }
  return &NoteRepository{}, fmt.Errorf("Unknow database connector")
}

func (nr *NoteRepository) CreateNote(author, title, content string) (n e.Note, err error) {
  return e.Note{}, fmt.Errorf("Something goes wrong!")
}

func (nr *NoteRepository) GetNoteByTitle(title string) (n []e.Note, err error) { return }

func (nr *NoteRepository) GetNotesByAuthor(author string) (n []e.Note, err error) {
  db, err := nr.PG.Connect()
  if err != nil {
    return []e.Note{}, err
  }
  defer nr.PG.Close(db)
  notes := []e.Note{}
  r := db.Where("author = ?", author).Find(&notes)
  return notes, r.Error
}

func (nr *NoteRepository) GetNoteByAuthorAndTitle(author, title string) (n []e.Note, err error) { return }
