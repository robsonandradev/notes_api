package main

import (
	"fmt"
	"github.com/gorilla/mux"
	hc "github.com/robsonandradev/notes_api/use_cases/health"
	"github.com/robsonandradev/notes_api/use_cases/login"
	"github.com/robsonandradev/notes_api/use_cases/read_note"
	"github.com/robsonandradev/notes_api/use_cases/create_user"
	"github.com/robsonandradev/notes_api/use_cases/read_user"
	"log"
	"net/http"
)

func main() {
	port := 3000
	router := mux.NewRouter()
	l := &login.LoginController{}
	setControllers(router, l)
	router.Use(logMiddleware, headerMiddleware, l.AuthenticationMiddleware)
	svc := &http.Server{
    Handler: router,
		Addr:    fmt.Sprintf("0.0.0.0:%d", port),
	}
	log.Println(fmt.Sprintf("magic happens on %d", port))
	log.Fatal(svc.ListenAndServe())
}

func setControllers(r *mux.Router, l *login.LoginController) {
	hc.Set(r)
	l.Set(r)
  readnote.NewReadNoteController().Set(r)
  createuser.NewCreateUserServiceController().Set(r)
  readuser.NewReadUserController().Set(r)
}

func logMiddleware(next http.Handler) http.Handler {
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
