package entities

import "time"

type Note struct {
  Author  string
  Title   string
  Content string
  Creation time.Time
  LastUpdate time.Time
}

func NewNote(author, title, content string, creation, lastUpdate time.Time) (n Note) {
  return Note{
    Author:     author,
    Title:      title,
    Content:    content,
    Creation:   creation,
    LastUpdate: lastUpdate,
  }
}
