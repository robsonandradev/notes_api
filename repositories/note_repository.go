package repositories

import (
	"fmt"

	"github.com/robsonandradev/notes_api/config"
	e "github.com/robsonandradev/notes_api/entities"
)

type INoteRepository interface {
	CreateNote(author, title, content string) (n e.Note, err error)
	GetNoteByTitle(title string) (notes []e.Note, err error)
	GetNotesByAuthor(author string) (notes []e.Note, err error)
	GetNoteByAuthorAndTitle(author, title string) (notes []e.Note, err error)
}

type NoteRepository struct {
	PG DBConnection
}

func NewNoteRepository(connector string) (*NoteRepository, error) {
	consts := config.NewConstants()
	if connector == consts.POSTGRES {
		ur := &NoteRepository{PG: &PostgresCon{}}
		return ur, nil
	}
	errorMsgs := config.NewErrorMessages()
	return &NoteRepository{}, fmt.Errorf(errorMsgs.UNKNOW_DATABASE_CONNECTOR)
}

func (nr *NoteRepository) CreateNote(author, title, content string) (n e.Note, err error) {
	return e.Note{}, fmt.Errorf("Something goes wrong!")
}

func (nr *NoteRepository) GetNoteByTitle(title string) (notes []e.Note, err error) {
	db, err := nr.PG.Connect()
	if err != nil {
		return
	}
	defer nr.PG.Close(db)
	r := db.Where("title like ?", "%"+title+"%").Find(&notes)
	return notes, r.Error
}

func (nr *NoteRepository) GetNotesByAuthor(author string) (notes []e.Note, err error) {
	db, err := nr.PG.Connect()
	if err != nil {
		return
	}
	defer nr.PG.Close(db)
	r := db.Where("author like ?", "%"+author+"%").Find(&notes)
	return notes, r.Error
}

func (nr *NoteRepository) GetNoteByAuthorAndTitle(author, title string) (notes []e.Note, err error) {
	db, err := nr.PG.Connect()
	if err != nil {
		return
	}
	defer nr.PG.Close(db)
  author = "%"+author+"%"
  title  = "%"+title+"%"
	r := db.Where("author like ? and title like ?", author, title).Find(&notes)
	return notes, r.Error
}
