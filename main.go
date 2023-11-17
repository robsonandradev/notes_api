package main

import (
  "net/http"
  "log"
  "fmt"
  "strings"
  "encoding/json"
  "github.com/gorilla/mux"
  "github.com/golang-jwt/jwt/v5"
  hc "github.com/robsonandradev/notes_api/use_cases/health"
  "github.com/robsonandradev/notes_api/use_cases/login"
)

func main() {
  port := 3000
  router := mux.NewRouter()
  setControllers(router)
  router.Use(loggingMiddleware, headerMiddleware, authenticationMiddleware)
  svc := &http.Server{
    Handler: router,
    Addr: fmt.Sprintf("0.0.0.0:%d", port),
  }
  log.Println(fmt.Sprintf("magic happens on %d", port))
  log.Fatal(svc.ListenAndServe())
}

func setControllers(r *mux.Router) {
  hc.Set(r)
  login.Set(r)
}

func loggingMiddleware(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    log.Println(r.RequestURI)
    next.ServeHTTP(w, r)
  })
}

func headerMiddleware(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    next.ServeHTTP(w, r)
  })
}

func authenticationMiddleware(next http.Handler) http.Handler {
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
      if !t.Valid {
        return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
      }
      return []byte("mysecret"), nil
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
