package login

import (
  "slices"
  "fmt"
  repos "github.com/robsonandradev/notes_api/repositories"
  e "github.com/robsonandradev/notes_api/entities"
)

type Login struct {}

func New() Login {
  return Login{}
}

func (l Login) doLogin(username, password string) (user e.User, err error) {
  userRepository := repos.UserRepository{}
  user, err = userRepository.getUserByUsernameAndPassword(username, password)
  return
}

/*
func (l Login) userExists(username string) (bool, error) {
  if !slices.Contains(usernames, username) {
    return false, fmt.Errorf("user not found!")
  }
  return true, nil
}

func (l Login) passwordMatches(username, password string) (bool, error) {
  return true, nil
}
*/
