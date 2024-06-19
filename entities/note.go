package entities

import (
  "time"
)

type Note struct {
  Id      string `gorm:"primaryKey"`
  Author  string
  Title   string
  Content string
  Creation time.Time
  LastUpdate time.Time
}

func NewNote(id, author, title, content string, creation, lastUpdate time.Time) (n Note) {
  return Note{
    Id:         id,
    Author:     author,
    Title:      title,
    Content:    content,
    Creation:   creation,
    LastUpdate: lastUpdate,
  }
}
