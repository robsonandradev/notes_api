package repositories

import (
  "fmt"
  "log"
  e "github.com/robsonandradev/notes_api/entities"
)

type IUserRepository interface {
  GetUserByUsernameAndPassword(username, password string) (e.User, error)
  GetUserByUsername(username string) (e.User, error)
}

type UserRepository struct {
  PG DBConnection
}

func NewUserRepository(connector string) (*UserRepository, error) {
  if connector == "postgres" {
    ur := &UserRepository{ PG: &PostgresCon{} }
    return ur, nil
  }
  return &UserRepository{}, fmt.Errorf("Unknow database connector")
}

func (ur *UserRepository) GetUserByUsername(username string) (e.User, error) {
  db, err := ur.PG.Connect()
  if err != nil {
    return e.User{}, err
  }
  defer ur.PG.Close(db)
  user := e.User{}
  r := db.First(&user, "username = ?", username)
  if r.Error != nil { log.Println("User not found!") }
  return user, r.Error
}

func (ur *UserRepository) GetUserByUsernameAndPassword(username, password string) (e.User, error) {
  db, err := ur.PG.Connect()
  if err != nil {
    return e.User{}, err
  }
  defer ur.PG.Close(db)
  user := e.User{}
  r := db.First(&user, "username = ? and password = ?", username, password)
  if r.Error != nil { log.Println("Wrong password!") }
  return user, r.Error
}
