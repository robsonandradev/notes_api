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
  return e.User{}, fmt.Errorf("user not found!")
}

// TODO: Implements database query
func (ur *UserRepository) GetUserByUsernameAndPassword(username, password string) (e.User, error) {
  return e.User{}, fmt.Errorf("wrong password!")
}
