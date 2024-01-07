package readnote

import (
  "fmt"
  "strings"
  e "github.com/robsonandradev/notes_api/entities"
  nr "github.com/robsonandradev/notes_api/repositories"
  "github.com/robsonandradev/notes_api/config"
)

type ReadNoteService struct {
  errorMsgs      config.ErrorMessages
  noteRepository nr.INoteRepository
}

func NewReadNoteService(noteRepo nr.INoteRepository) *ReadNoteService {
  return &ReadNoteService{
    noteRepository: noteRepo,
    errorMsgs: *config.NewErrorMessages(),
  }
}

func (rn *ReadNoteService) GetNoteByTitle(title string) ([]e.Note, error) {
  if strings.TrimSpace(title) == "" {
    return []e.Note{}, fmt.Errorf(rn.errorMsgs.FIELD_TITLE_SHOULD_NOT_BE_EMPTY)
  }
  return rn.noteRepository.GetNoteByTitle(title)
}

func (rn *ReadNoteService) GetNoteByAuthorAndTitle(author, title string) ([]e.Note, error) {
  if strings.TrimSpace(author) == "" && strings.TrimSpace(title) == "" {
    return []e.Note{}, fmt.Errorf(rn.errorMsgs.FIELD_AUTHOR_AND_TITLE_SHOULD_NOT_BE_EMPTY)
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
    return []e.Note{}, fmt.Errorf(rn.errorMsgs.FIELD_AUTHOR_SHOULD_NOT_BE_EMPTY)
  }
  return rn.noteRepository.GetNotesByAuthor(author)
}
