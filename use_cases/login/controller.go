package login

import (
  "time"
  "io"
  "encoding/json"
  "net/http"
  "log"
  "strings"
  "github.com/gorilla/mux"
  "github.com/golang-jwt/jwt/v5"
  repos "github.com/robsonandradev/notes_api/repositories"
  "github.com/robsonandradev/notes_api/config"
)

type responseToken struct {
  Token string `json:"token"`
}

type requestBody struct {
  Username string `json:"username"`
  Password string `json:"password"`
}

type LoginController struct {
  jwtClaims *jwt.RegisteredClaims
}

func (l *LoginController) Set(router *mux.Router) {
  router.HandleFunc("/login", l.exec).Methods("POST")
}

func (l *LoginController) exec(w http.ResponseWriter, r *http.Request) {
  consts := config.NewConstants()
  repo, err := repos.NewUserRepository(consts.POSTGRES)
  if err != nil { panic(err) }
  login := New(repo)
  u := getRequestBody(r.Body)
  user, err := login.doLogin(u.Username, u.Password)
  if err != nil {
    w.WriteHeader(http.StatusUnauthorized)
    setErrors(w, err.Error())
    return
  }
  tokenString, err := l.signin(user.Username)
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

func (l *LoginController) signin(u string) (string, error) {
  // TODO: Figure out which data will be a good fit for this token
	expirationTime := time.Now().Add(5 * time.Minute)
  claims := &jwt.RegisteredClaims{
    ExpiresAt: jwt.NewNumericDate(expirationTime),
    Issuer: "Note API",
    Subject: u,
  }
  l.jwtClaims = claims
  log.Println(l.jwtClaims)
  token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
  // TODO: Receive the secret by environment variable
  return token.SignedString([]byte("mysecret"))
}

func (l *LoginController) AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    if r.RequestURI == "/login" {
      next.ServeHTTP(w, r)
      return
    }
		authArr := strings.Split(r.Header.Get("Authorization"), " ")
    errorMsg := "user is not logged in"
    if len(authArr) < 2 {
      w.WriteHeader(http.StatusUnauthorized)
      setErrors(w, errorMsg)
      return
    }
    tokenString := authArr[1]
    parser := jwt.Parser{}
    expirationTime := time.Now().Add(5 * time.Minute)
    if l.jwtClaims == nil {
      w.WriteHeader(http.StatusUnauthorized)
      log.Println("Empty jwtClaims!")
      setErrors(w, errorMsg)
      return
    }
    l.jwtClaims.ExpiresAt = jwt.NewNumericDate(expirationTime)
    _, err := parser.ParseWithClaims(tokenString, l.jwtClaims, func(t *jwt.Token) (interface{}, error) {
      // TODO: Receive the secret by environment variable
      mySigningKey := []byte("mysecret")
      return mySigningKey, nil
    })
    if err != nil {
      w.WriteHeader(http.StatusUnauthorized)
      log.Println(err)
      setErrors(w, errorMsg)
      return
    }
    next.ServeHTTP(w, r)
	})
}
