package login

import (
  "log"
  repos "github.com/robsonandradev/notes_api/repositories"
  e "github.com/robsonandradev/notes_api/entities"
)

type loginSvc struct {
  userRepository repos.IUserRepository
}

func New(ur repos.IUserRepository) loginSvc {
  return loginSvc{ userRepository: ur }
}

func (l loginSvc) doLogin(username, password string) (user e.User, err error) {
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
