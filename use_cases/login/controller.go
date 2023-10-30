package login

import (
  "time"
  "io"
  "encoding/json"
  "net/http"
  "github.com/gorilla/mux"
  "github.com/golang-jwt/jwt/v5"
  repos "github.com/robsonandradev/notes_api/repositories"
)

type responseToken struct {
  Token string `json:"token"`
}

type requestBody struct {
  Username string `json:"username"`
  Password string `json:"password"`
}

func Set(router *mux.Router) {
  router.HandleFunc("/login", exec).Methods("POST")
}

func exec(w http.ResponseWriter, r *http.Request) {
  login := New(&repos.UserRepository{})
  u := getRequestBody(r.Body)
  user, err := login.doLogin(u.Username, u.Password)
  if err != nil {
    w.WriteHeader(http.StatusUnauthorized)
    json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
    return
  }
  w.WriteHeader(http.StatusOK)
  tokenString, err := getToken(user.Username)
  if err != nil { panic(err) }
  json.NewEncoder(w).Encode(responseToken{Token: tokenString})
}

func getRequestBody(b io.ReadCloser) *requestBody {
  var u requestBody
  err := json.NewDecoder(b).Decode(&u)
  if err != nil { panic(err) }
  return &u
}

func getToken(u string) (string, error) {
  // TODO: Figure out which data will be a good fit for this token
  token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
    "username": u,
    "nbf": time.Now().UTC().Unix(),
  })
  // TODO: Figure out a better way to set this secret
  secret, err := jwt.SigningMethodHS256.Sign("root", []byte("mysecret"))
  if err != nil {
    return "", err
  }
  return token.SignedString(secret)
}
