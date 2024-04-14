package createuser

import (
  //"time"
  "io"
  "encoding/json"
  "net/http"
  //"log"
  //"strings"
  "github.com/gorilla/mux"
  repos "github.com/robsonandradev/notes_api/repositories"
  "github.com/robsonandradev/notes_api/config"
)

type responseUser struct {
  Id       string `json:"id"`
  Username string `json:"username"`
  Email    string `json:"emal"`
}

type requestBody struct {
  Username string `json:"username"`
  Password string `json:"password"`
  Email    string `json:"email"`
}

type CreateUserController struct {}

func NewCreateUserServiceController() *CreateUserController {
  return &CreateUserController{}
}

func (c *CreateUserController) Set(router *mux.Router) {
  router.HandleFunc("/user", c.run).Methods("PUT")
}

func (c *CreateUserController) run(w http.ResponseWriter, r *http.Request) {
  consts := config.NewConstants()
  repo, err := repos.NewUserRepository(consts.POSTGRES)
  if err != nil { panic(err) }
  CreateUserSvc := NewCreateUserService(repo)
  u := getRequestBody(r.Body)
  user, err := CreateUserSvc.Run(u.Username, u.Password, u.Email)
  if err != nil {
    w.WriteHeader(http.StatusBadRequest)
    setErrors(w, err.Error())
    return
  }
  w.WriteHeader(http.StatusOK)
  json.NewEncoder(w).Encode(responseUser{
    Id: user.Id,
    Username: user.Username,
    Email: user.Email,
  })
}

func setErrors(w http.ResponseWriter, err string) {
  json.NewEncoder(w).Encode(map[string]string{"error": err})
}

func getRequestBody(b io.ReadCloser) *requestBody {
  var u requestBody
  err := json.NewDecoder(b).Decode(&u)
  if err != nil { panic(err) }
  return &u
}
