package createnote

import (
  "os"
  "time"
  "fmt"
  "testing"
  "reflect"
  e "github.com/robsonandradev/notes_api/entities"
)

var (
  note e.Note
  createNote CreateNote
)


func TestMain(m *testing.M) {
  fmt.Println('T')
  now := time.Now().UTC()
  note = e.NewNote("john", "My Note", "Loren ipson", now, now)
  repo := NoteRepositoryMock{}
  createNote = New(&repo)
  os.Exit(m.Run())
}

func TestSuccessfulNoteCreation(t *testing.T) {
  t.Run("when note metadata was set should create a new note", func(t *testing.T) {
    want := note
    have, err := createNote.Run("john wick", "My Note", "loren ipson")
    if err != nil {
      panic(err)
    }
    if reflect.DeepEqual(want, have) {
      t.Errorf("want %s, and have %s", want, have)
    }
  })
}

func TestCreateNoteEmptyAuthor(t *testing.T) {
  t.Run("when author is empty should return an error", func(t *testing.T) {
    want := fmt.Errorf("Author field shouldn't be empty")
    _, have := createNote.Run("", "My Note", "login ipson")
    if want.Error() != have.Error() {
      t.Errorf("want %s, and have %s", want, have)
    }
  })
}

func TestCreateNoteEmptyTitle(t *testing.T) {
  t.Run("when title is empty should return an error", func(t *testing.T) {
    want := fmt.Errorf("Title field shouldn't be empty")
    _, have := createNote.Run("John Wick", "", "login ipson")
    if want.Error() != have.Error() {
      t.Errorf("want %s, and have %s", want, have)
    }
  })
}

func TestCreateNoteEmptyContent(t *testing.T) {
  t.Run("when content is empty should return an error", func(t *testing.T) {
    want := fmt.Errorf("Content field shouldn't be empty")
    _, have := createNote.Run("John Wick", "My Note", "")
    if want.Error() != have.Error() {
      t.Errorf("want %s, and have %s", want, have)
    }
  })
}

// starting mocks
type NoteRepositoryMock struct {}

func (nr *NoteRepositoryMock) CreateNote(author, title, content string) (n e.Note, err error) {
  note.Author  = author
  note.Title   = title
  note.Content = content
  return note, nil
}

func (NoteRepositoryMock) GetNoteByTitle(title string) (n e.Note, err error) {return}
func (NoteRepositoryMock) GetNotesByAuthor(author string) (n e.Note, err error) {return}
func (NoteRepositoryMock) GetNoteByAuthorAndTitle(author, title string) (n e.Note, err error) {return}
