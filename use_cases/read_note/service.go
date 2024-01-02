package readnote

import (
  "fmt"
  "strings"
  e "github.com/robsonandradev/notes_api/entities"
  nr "github.com/robsonandradev/notes_api/repositories"
)

type ReadNoteService struct {
  noteRepository nr.INoteRepository
}

func NewReadNoteService(noteRepo nr.INoteRepository) *ReadNoteService {
  return &ReadNoteService{noteRepository: noteRepo}
}

func (rn *ReadNoteService) GetNoteByTitle(title string) ([]e.Note, error) {
  if strings.TrimSpace(title) == "" {
    return []e.Note{}, fmt.Errorf("Field title should not be empty!")
  }
  return rn.noteRepository.GetNoteByTitle(title)
}

func (rn *ReadNoteService) GetNoteByAuthorAndTitle(author, title string) ([]e.Note, error) {
  if strings.TrimSpace(author) == "" && strings.TrimSpace(title) == "" {
    return []e.Note{}, fmt.Errorf("Field author and title should not be empty!")
  }
  if strings.TrimSpace(author) == "" && strings.TrimSpace(title) != "" {
    return rn.GetNoteByTitle(title)
  }
  if strings.TrimSpace(author) != "" && strings.TrimSpace(title) == "" {
    return rn.GetNotesByAuthor(author)
  }
  return rn.noteRepository.GetNoteByAuthorAndTitle(author, title)
}

func (rn *ReadNoteService) GetNotesByAuthor(author string) ([]e.Note, error) {
  if strings.TrimSpace(author) == "" {
    return []e.Note{}, fmt.Errorf("Field author should not be empty!")
  }
  return rn.noteRepository.GetNotesByAuthor(author)
}
