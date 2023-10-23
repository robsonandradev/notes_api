package main

import (
  "net/http"
  "log"
  "fmt"
  "github.com/gorilla/mux"
  hc "github.com/robsonandradev/notes_api/use_cases/health"
)

func main() {
  port := 3000
  router := mux.NewRouter()
  hc.Set(router)
  router.Use(loggingMiddleware, headerMiddleware)
  svc := &http.Server{
    Handler: router,
    Addr: fmt.Sprintf("0.0.0.0:%d", port),
  }
  log.Println(fmt.Sprintf("magic happens on %d", port))
  log.Fatal(svc.ListenAndServe())
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
