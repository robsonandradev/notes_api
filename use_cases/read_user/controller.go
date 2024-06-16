package readuser

import (
  "log"
  "encoding/json"
  "net/http"
  "github.com/gorilla/mux"
  repos "github.com/robsonandradev/notes_api/repositories"
  "github.com/robsonandradev/notes_api/config"
)

type ReadUserController struct {
  errorMsgs *config.ErrorMessages
}

func NewReadUserController() *ReadUserController {
  return &ReadUserController{errorMsgs: config.NewErrorMessages()}
}

func (c *ReadUserController) Set(router *mux.Router) {
  router.HandleFunc("/users", c.exec).Methods("GET")
}

func (c *ReadUserController) exec(w http.ResponseWriter, r *http.Request) {
  consts := config.NewConstants()
  repo, err := repos.NewUserRepository(consts.POSTGRES)
  if err != nil {
    internalServerError(w, err)
    return
  }
  readUserSvc := NewReadUserSvc(repo)
  username := r.URL.Query().Get("username")
  email    := r.URL.Query().Get("email")
  id       := r.URL.Query().Get("id")
  notes, err := readUserSvc.Run(username, email, id)
  if err != nil {
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
