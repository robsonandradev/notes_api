package login

import (
  "slices"
  "fmt"
)

type Login struct {}

func New() Login {
  return Login{}
}

func (l Login) doLogin(username, password string) (result bool, err error) {
  result, err = l.userExists(username)
  if err != nil { return }
  result, err = l.passwordMatches(username, password)
  if err != nil { return }
  return
}

func (l Login) userExists(username string) (bool, error) {
  usernames := []string{"john"}
  if !slices.Contains(usernames, username) {
    return false, fmt.Errorf("user not found!")
  }
  return true, nil
}

func (l Login) passwordMatches(username, password string) (bool, error) {
  return true, nil
}
