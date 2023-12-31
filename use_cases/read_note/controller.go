package readnote

import (
  "encoding/json"
  "net/http"
  "github.com/gorilla/mux"
  repos "github.com/robsonandradev/notes_api/repositories"
)

type ReadNoteController struct {}

func (c *ReadNoteController) Set(router *mux.Router) {
  router.HandleFunc("/notes", c.exec).Methods("GET")
}

func (c *ReadNoteController) exec(w http.ResponseWriter, r *http.Request) {
  repo, err := repos.NewNoteRepository("postgres")
  if err != nil { panic(err) }
  readNoteSvc := NewReadNoteService(repo)
  author := r.URL.Query().Get("author")
  title  := r.URL.Query().Get("title")
  notes, err := readNoteSvc.GetNoteByAuthorAndTitle(author, title)
  if err != nil { panic(err) }
  w.WriteHeader(http.StatusOK)
  json.NewEncoder(w).Encode(notes)
}
