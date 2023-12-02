package login

import (
  "time"
  "io"
  "encoding/json"
  "net/http"
  "log"
  "fmt"
  "strings"
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

type Login struct {}

func (l Login) Set(router *mux.Router) {
  router.HandleFunc("/login", exec).Methods("POST")
}

func exec(w http.ResponseWriter, r *http.Request) {
  repo, err := repos.NewUserRepository("postgres")
  if err != nil { panic(err) }
  login := New(repo)
  u := getRequestBody(r.Body)
  user, err := login.doLogin(u.Username, u.Password)
  if err != nil {
    w.WriteHeader(http.StatusUnauthorized)
    //json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
    setErrors(w, err.Error())
    return
  }
  tokenString, err := signin(user.Username)
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    setErrors(w, err.Error())
    return
  }
  w.WriteHeader(http.StatusOK)
  json.NewEncoder(w).Encode(responseToken{Token: tokenString})
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

func signin(u string) (string, error) {
  // TODO: Figure out which data will be a good fit for this token
	expirationTime := time.Now().Add(5 * time.Minute)
  token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
    "username": u,
    "exp": expirationTime.Unix(),
  })
  // TODO: Receive the secret by environment variable
  secret, err := jwt.SigningMethodHS256.Sign("root", []byte("mysecret"))
  if err != nil {
    return "", err
  }
  return token.SignedString(secret)
}

func (l Login) AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    if r.RequestURI == "/login" {
      next.ServeHTTP(w, r)
      return
    }
		authArr := strings.Split(r.Header.Get("Authorization"), " ")
    errorMsg := map[string]string{"error": "user is not logged in"}
    if len(authArr) < 2 {
      w.WriteHeader(http.StatusUnauthorized)
      json.NewEncoder(w).Encode(errorMsg)
      return
    }
    tokenString := authArr[1]
    log.Println(tokenString)
    parser := jwt.Parser{}
    token, err := parser.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
      //if _, ok := t.Method.(*jwt.SigningMethodHS256); !ok {
      log.Println(t)
      if !t.Valid {
        return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
      }
      return tokenString, nil
    })
    if err != nil {
      w.WriteHeader(http.StatusUnauthorized)
      log.Println(err)
      json.NewEncoder(w).Encode(errorMsg)
      return
    }
    log.Println(token.Valid)
    next.ServeHTTP(w, r)
	})
}
