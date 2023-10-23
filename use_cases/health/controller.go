package health

import (
  "encoding/json"
  "net/http"
  "github.com/gorilla/mux"
)

type ping struct {
  Ping string `json:"ping"`
}

func Set(router *mux.Router) {
  router.HandleFunc("/", check).Methods("GET")
  router.HandleFunc("/health", check).Methods("GET")
}

func check(w http.ResponseWriter, r *http.Request) {
  w.WriteHeader(http.StatusOK)
  out := ping{Ping: "pong"}
  err := json.NewEncoder(w).Encode(out)
  if err != nil { panic(err) }
  //outJson, err := json.Marshal(out)
  //if err != nil { panic(err) }
  //log.Println("handling request", string(outJson))
}

