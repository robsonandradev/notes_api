package createuser

import (
  "log"
  "fmt"
  repos "github.com/robsonandradev/notes_api/repositories"
  e "github.com/robsonandradev/notes_api/entities"
)

type CreateUserSvc struct {
  userRepository repos.IUserRepository
}

func NewCreateUserService(ur repos.IUserRepository) CreateUserSvc {
  return CreateUserSvc{ userRepository: ur }
}

func (cr CreateUserSvc) Run(username, password, email string) (user e.User, err error) {
  log.Println("Register new user:", username)
  if username == "" {
    return e.User{}, fmt.Errorf("Username is required!")
  }
  if password == "" {
    return e.User{}, fmt.Errorf("Password is required!")
  }
  if email == "" {
    return e.User{}, fmt.Errorf("Email is required!")
  }
  user, err = cr.userRepository.CreateUser(username, password, email)
  return
}
