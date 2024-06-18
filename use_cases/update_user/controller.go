package updateuser

import (
  "log"
  "encoding/json"
  "net/http"
  "github.com/gorilla/mux"
  repos "github.com/robsonandradev/notes_api/repositories"
  "github.com/robsonandradev/notes_api/config"
)

type UpdateUserController struct {
  errorMsgs *config.ErrorMessages
}

type requestBody struct {
  Id       string `json:"id"`
  Email    string `json:"email"`
  Password string `json:"password"`
}

func NewUpdateUserController() *UpdateUserController {
  return &UpdateUserController{errorMsgs: config.NewErrorMessages()}
}

func (c *UpdateUserController) Set(router *mux.Router) {
  router.HandleFunc("/user", c.exec).Methods("POST")
}

func (c *UpdateUserController) exec(w http.ResponseWriter, r *http.Request) {
  consts := config.NewConstants()
  repo, err := repos.NewUserRepository(consts.POSTGRES)
  if err != nil {
    internalServerError(w, err)
    return
  }
  updateUserSvc := NewUpdateUserSvc(repo)
  var rb requestBody
  err = json.NewDecoder(r.Body).Decode(&rb)
  if err != nil { panic(err) }
  user, err := updateUserSvc.Run(rb.Id, rb.Password, rb.Email)
  if err != nil {
    if err.Error() == c.errorMsgs.USER_NOT_FOUND {
      returnError(w, err)
      return
    }
    if err.Error() == c.errorMsgs.USER_EMAIL_OR_PASSWORD_REQUIRED {
      returnError(w, err)
      return
    }
    internalServerError(w, err)
    return
  }
  w.WriteHeader(http.StatusOK)
  json.NewEncoder(w).Encode(user)
}

func internalServerError(w http.ResponseWriter, err error) {
  log.Println(err)
  w.WriteHeader(http.StatusInternalServerError)
  json.NewEncoder(w).Encode(map[string]string{"error": "internal server error"})
}

func returnError(w http.ResponseWriter, err error) {
  log.Println(err)
  w.WriteHeader(http.StatusForbidden)
  json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
}
