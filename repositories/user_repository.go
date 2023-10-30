package repositories

import (
  "fmt"
  e "github.com/robsonandradev/notes_api/entities"
)

type IUserRepository interface {
  GetUserByUsernameAndPassword(username, password string) (e.User, error)
  GetUserByUsername(username string) (e.User, error)
}

type UserRepository struct {}

// TODO: Implements database query
func (ur *UserRepository) GetUserByUsername(username string) (e.User, error) {
  if username == "john.wick@gmail.com" {
    return e.NewUser("1", username, "123", username), nil
  }
  return e.User{}, fmt.Errorf("user not found!")
}

// TODO: Implements database query
func (ur *UserRepository) GetUserByUsernameAndPassword(username, password string) (e.User, error) {
  if username == "john.wick@gmail.com" && password == "john.wick" {
    return e.NewUser("1", username, password, username), nil
  }
  return e.User{}, fmt.Errorf("wrong password!")
}
