package main

import (
  "log"
  "gorm.io/gorm"
  "github.com/google/uuid"
  repos "github.com/robsonandradev/notes_api/repositories"
  e "github.com/robsonandradev/notes_api/entities"
)

func main() {
  log.Printf("starting db migration")
  pg := repos.PostgresCon{}
  db, err := pg.Connect()
  if err != nil {
    panic(err)
  }
  defer pg.Close(db)
  log.Printf("creating tables")
  db.AutoMigrate(&e.Note{}, &e.User{})
  log.Printf("starting seed")
  //dbSeed(db)
  log.Printf("finished successfuly")
}

func dbSeed(db *gorm.DB) {
  r := db.Create(e.User{
    Id: uuid.NewString(), 
    Email: "john.wick@gmail.com",
    Username: "john.wick@gmail.com",
    Password: "john.wick",
  })
  if r.Error != nil {
    panic(r.Error)
  }
  log.Printf("User created")
}
