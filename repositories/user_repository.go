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

func (UserRepository) GetUserByUsername(username string) (e.User, error) {
  users := mockUsers()
  for _, user := range users {
    if user.Username == username {
      return user, nil
    }
  }
  return e.User{}, fmt.Errorf("user not found!")
}

func (ur UserRepository) GetUserByUsernameAndPassword(username, password string) (e.User, error) {
  users := mockUsers()
  for _, user := range users {
    if user.Username == username && user.Password == password {
        return user, nil
    }
  }
  return e.User{}, fmt.Errorf("wrong password!")
}

func mockUsers() []e.User {
  u1 := e.NewUser("1", "john", "john", "john@gmail.com")
  u2 := e.NewUser("2", "jone", "jone", "jone@gmail.com")
  u3 := e.NewUser("3", "joel", "joel", "joel@gmail.com")
  return []e.User{u1, u2, u3}
}
