package login

import (
  repos "github.com/robsonandradev/notes_api/repositories"
  e "github.com/robsonandradev/notes_api/entities"
)

type Login struct {}

func New() Login {
  return Login{}
}

func (l Login) doLogin(username, password string) (user e.User, err error) {
  userRepository := repos.UserRepository{}
  user, err = userRepository.GetUserByUsername(username)
  if err != nil {return}
  user, err = userRepository.GetUserByUsernameAndPassword(username, password)
  return
}
