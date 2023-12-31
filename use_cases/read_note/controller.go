package readnote

import (
  "log"
  "encoding/json"
  "net/http"
  "github.com/gorilla/mux"
  repos "github.com/robsonandradev/notes_api/repositories"
  "github.com/robsonandradev/notes_api/config"
)

type ReadNoteController struct {
  errorMsgs *config.ErrorMessages
}

func NewReadNoteController() *ReadNoteController {
  return &ReadNoteController{errorMsgs: config.NewErrorMessages()}
}

func (c *ReadNoteController) Set(router *mux.Router) {
  router.HandleFunc("/notes", c.exec).Methods("GET")
}

func (c *ReadNoteController) exec(w http.ResponseWriter, r *http.Request) {
  consts := config.NewConstants()
  repo, err := repos.NewNoteRepository(consts.POSTGRES)
  if err != nil {
    internalServerError(w, err)
    return
  }
  readNoteSvc := NewReadNoteService(repo)
  author := r.URL.Query().Get("author")
  title  := r.URL.Query().Get("title")
  notes, err := readNoteSvc.GetNoteByAuthorAndTitle(author, title)
  if err != nil {
    if err.Error() == c.errorMsgs.FIELD_AUTHOR_AND_TITLE_SHOULD_NOT_BE_EMPTY {
      returnError(w, err)
      return
    }
    internalServerError(w, err)
    return
  }
  w.WriteHeader(http.StatusOK)
  json.NewEncoder(w).Encode(notes)
}

func internalServerError(w http.ResponseWriter, err error) {
  log.Println(err)
  w.WriteHeader(http.StatusInternalServerError)
  json.NewEncoder(w).Encode(map[string]string{"error": "internal server error"})
}

func returnError(w http.ResponseWriter, err error) {
  log.Println(err)
  w.WriteHeader(http.StatusNotFound)
  json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
}
