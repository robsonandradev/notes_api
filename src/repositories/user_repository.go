package repositories

import (
  e "github.com/robsonandradev/src/entities"
)

type IUserRepository interface {
  getUserByUsernameAndPassword(username, password string) (e.User, error)
}

type UserRepository struct {}

func (ur UserRepository) getUserByUsernameAndPassword(username, password string) (e.User, error) {
  user := e.User.New()
  return user, nil
}
