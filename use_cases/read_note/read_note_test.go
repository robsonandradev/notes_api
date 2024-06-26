package readnote

import (
	"fmt"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/robsonandradev/notes_api/config"
	e "github.com/robsonandradev/notes_api/entities"
)

var (
	note            e.Note
	readNoteService *ReadNoteService
	errorMsgs       *config.ErrorMessages
)

func TestMain(m *testing.M) {
	now := time.Now()
	note = e.NewNote("1", "john wick", "my note", "loren ipson and go on", now, now)
	noteRepo := NoteRepositoryMock{}
	readNoteService = NewReadNoteService(&noteRepo)
	errorMsgs = config.NewErrorMessages()
	os.Exit(m.Run())
}

func TestSuccessfulGetNote(t *testing.T) {
	t.Run("when search for note by title and note exists then return the note", func(t *testing.T) {
		want := []e.Note{note}
		have, err := readNoteService.GetNoteByTitle("my note")
		if err != nil {
			panic(err)
		}
		if !reflect.DeepEqual(want, have) {
			t.Errorf("want %s, and have %s", want, have)
		}
	})

	t.Run("when search for notes by author and note exists then return the note", func(t *testing.T) {
		want := []e.Note{note}
		have, err := readNoteService.GetNotesByAuthor("john wick")
		if err != nil {
			panic(err)
		}
		if !reflect.DeepEqual(want, have) {
			t.Errorf("want %s, and have %s", want, have)
		}
	})

	t.Run("when search for note by author and title and note exists then return the note", func(t *testing.T) {
		want := []e.Note{note}
		have, err := readNoteService.GetNoteByAuthorAndTitle("john wick", "my note")
		if err != nil {
			panic(err)
		}
		if !reflect.DeepEqual(want, have) {
			t.Errorf("want %s, and have %s", want, have)
		}
	})

	t.Run("when search for a existent note by author and title and author field is empty then return note", func(t *testing.T) {
		want := []e.Note{note}
		have, err := readNoteService.GetNoteByAuthorAndTitle("", "my note")
		if err != nil {
			panic(err)
		}
		if !reflect.DeepEqual(want, have) {
			t.Errorf("want %s, but have %s", want, have)
		}
	})

	t.Run("when search for a existent note by author and title and title field is empty then return note", func(t *testing.T) {
		want := []e.Note{note}
		have, err := readNoteService.GetNoteByAuthorAndTitle("john wick", "")
		if err != nil {
			panic(err)
		}
		if !reflect.DeepEqual(want, have) {
			t.Errorf("want %s, but have %s", want, have)
		}
	})
}

func TestGetNoteWithEmptyField(t *testing.T) {
	t.Run("when search for note by title and field is empty then return an error", func(t *testing.T) {
		want := fmt.Errorf(errorMsgs.FIELD_TITLE_SHOULD_NOT_BE_EMPTY)
		_, have := readNoteService.GetNoteByTitle("")
		if have == nil || want.Error() != have.Error() {
			t.Errorf("want %s, but have %s", want, have)
		}
	})

	t.Run("when search for notes by author and field is empty then return an error", func(t *testing.T) {
		want := fmt.Errorf(errorMsgs.FIELD_AUTHOR_SHOULD_NOT_BE_EMPTY)
		_, have := readNoteService.GetNotesByAuthor("")
		if have == nil || want.Error() != have.Error() {
			t.Errorf("want %s, but have %s", want, have)
		}
	})

	t.Run("when search for note by author and title and fields are empty retun error", func(t *testing.T) {
		want := fmt.Errorf(errorMsgs.FIELD_AUTHOR_AND_TITLE_SHOULD_NOT_BE_EMPTY)
		_, have := readNoteService.GetNoteByAuthorAndTitle("", "")
		if have == nil || want.Error() != have.Error() {
			t.Errorf("want %s, but have %s", want, have)
		}
	})
}

func TestGetNoteWhichDoesntExist(t *testing.T) {
	t.Run("when search for note by title that doesnt exist then return an error", func(t *testing.T) {
    want := 0
		have, _ := readNoteService.GetNoteByTitle("note x")
    if len(have) != want {
			t.Errorf("want %d, but have %s", want, have)
		}
	})

	t.Run("when search for notes by author that doesnt exist then return an error", func(t *testing.T) {
    want := 0
		have, _ := readNoteService.GetNotesByAuthor("Elon Musk")
    if len(have) != want {
			t.Errorf("want %d, but have %s", want, have)
		}
	})

	t.Run("when search for note by AUTHOR and title that doesnt exist then return an error", func(t *testing.T) {
    want := 0
		have, _ := readNoteService.GetNoteByAuthorAndTitle("Elon Musk", "my note")
    if len(have) != want {
			t.Errorf("want %d, but have %s", want, have)
		}
	})

	t.Run("when search for note by author and TITLE that doesnt exist then return an error", func(t *testing.T) {
    want := 0
		have, _ := readNoteService.GetNoteByAuthorAndTitle("john wick", "xnote")
    if len(have) != want {
			t.Errorf("want %d, but have %s", want, have)
		}
	})
}

// starting mocks

type NoteRepositoryMock struct{}

func (nr *NoteRepositoryMock) CreateNote(author, title, content string) (n e.Note, err error) { return }
func (NoteRepositoryMock) GetNoteByTitle(title string) (notes []e.Note, err error) {
	if title != "my note" {
		return
	}
	notes = append(notes, note)
	return
}
func (NoteRepositoryMock) GetNotesByAuthor(author string) (notes []e.Note, err error) {
	if author != "john wick" {
		return
	}
	notes = append(notes, note)
	return
}
func (NoteRepositoryMock) GetNoteByAuthorAndTitle(author, title string) (notes []e.Note, err error) {
	if author == "john wick" && title == "my note" {
		notes = append(notes, note)
		return
	}
  return
}
