package readnote

import (
  "fmt"
  "strings"
  e "github.com/robsonandradev/notes_api/entities"
  nr "github.com/robsonandradev/notes_api/repositories"
)

type ReadNote struct {
  noteRepository nr.INoteRepository
}

func New(noteRepo nr.INoteRepository) *ReadNote {
  return &ReadNote{noteRepository: noteRepo}
}

// TODO: GetNote* should return array of notes
func (rn *ReadNote) GetNoteByTitle(title string) (e.Note, error) {
  if strings.TrimSpace(title) == "" {
    return e.Note{}, fmt.Errorf("Field title should not be empty!")
  }
  return rn.noteRepository.GetNoteByTitle(title)
}

func (rn *ReadNote) GetNoteByAuthorAndTitle(author, title string) (e.Note, error) {
  if strings.TrimSpace(author) == "" {
    return e.Note{}, fmt.Errorf("Field author should not be empty!")
  }
  if strings.TrimSpace(title) == "" {
    return e.Note{}, fmt.Errorf("Field title should not be empty!")
  }
  return rn.noteRepository.GetNoteByAuthorAndTitle(author, title)
}

func (rn *ReadNote) GetNotesByAuthor(author string) (e.Note, error) {
  if strings.TrimSpace(author) == "" {
    return e.Note{}, fmt.Errorf("Field author should not be empty!")
  }
  return rn.noteRepository.GetNotesByAuthor(author)
}
