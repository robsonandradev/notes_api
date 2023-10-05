package repositories

import (
  "fmt"
  e "github.com/robsonandradev/notes_api/entities"
)

type INoteRepository interface {
  CreateNote(author, title, content string) (n e.Note, err error)
  // TODO: GetNote* should return array of notes
  GetNoteByTitle(title string) (n e.Note, err error)
  GetNotesByAuthor(author string) (n e.Note, err error)
  GetNoteByAuthorAndTitle(author, title string) (n e.Note, err error)
}

type NoteRepository struct {}

func (nr *NoteRepository) CreateNote(author, title, content string) (n e.Note, err error) {
  return e.Note{}, fmt.Errorf("Something goes wrong!")
}
