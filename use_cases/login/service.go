package login

import (
  "log"
  repos "github.com/robsonandradev/notes_api/repositories"
  e "github.com/robsonandradev/notes_api/entities"
)

type Login struct {
  userRepository repos.IUserRepository
}

func New(ur repos.IUserRepository) Login {
  return Login{ userRepository: ur }
}

func (l Login) doLogin(username, password string) (user e.User, err error) {
  log.Println("logging with user:", username)
  user, err = l.userRepository.GetUserByUsername(username)
  if err != nil {
    log.Println("User not found!")
    return
  }
  user, err = l.userRepository.GetUserByUsernameAndPassword(username, password)
  if err != nil { log.Println("Wrong password!") }
  return
}
