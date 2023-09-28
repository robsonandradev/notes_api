package createnote

import (
  "fmt"
  e "github.com/robsonandradev/notes_api/entities"
  nr "github.com/robsonandradev/notes_api/repositories"
)

type CreateNote struct {
  NoteRepository nr.INoteRepository
}

func New(repo nr.INoteRepository) CreateNote {
  return CreateNote{NoteRepository: repo}
}

func (c *CreateNote) Run(author, title, content string) (n e.Note, err error) {
  if author == "" {
    return e.Note{}, fmt.Errorf("Author field shouldn't be empty")
  }
  if title == "" {
    return e.Note{}, fmt.Errorf("Title field shouldn't be empty")
  }
  if content == "" {
    return e.Note{}, fmt.Errorf("Content field shouldn't be empty")
  }
  return c.NoteRepository.CreateNote(author, title, content)
}
