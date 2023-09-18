package repositories

import (
  "fmt"
  e "github.com/robsonandradev/entities"
)

type IUserRepository interface {
  getUserByUsernameAndPassword(username, password string) (e.User, error)
}

type UserRepository struct {}

func (ur UserRepository) getUserByUsernameAndPassword(username, password string) (e.User, error) {
  users := mockUsers()
  isUserFound := false
  for _, user := range users {
    if user.username == username {
      isUserFound = true
      if user.password == password {
        return user, nil
      }
    }
  }
  if isUserFound {
    return nil, fmt.Errorf("user not found!")
  }
  return nil, fmt.Errorf("wrong password!")
}

func mockUsers() []e.User {
  u1 := e.User.New(1, "john", "john", "john@gmail.com")
  u2 := e.User.New(1, "jone", "jone", "jone@gmail.com")
  u3 := e.User.New(1, "joel", "joel", "joel@gmail.com")
  return []e.User{u1, u2, u3}
}
