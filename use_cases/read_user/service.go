package readuser

import (
  e "github.com/robsonandradev/notes_api/entities"
  repos "github.com/robsonandradev/notes_api/repositories"
)

type ReadUserSvc struct {
  userRepository repos.IUserRepository
}

func NewReadUserSvc(repo repos.IUserRepository) *ReadUserSvc {
  return &ReadUserSvc{userRepository: repo}
}

func (ru ReadUserSvc) Run(username, email, id string) (users []e.User, err error) {
  if (id == "" && username == "" && email == "") {
    users, err = ru.userRepository.GetAll()
    return
  }
  if (id != "") {
    u, err := ru.userRepository.GetUserById(id)
    return append(users, u), err
  }
  if (email != "") {
    u, err := ru.userRepository.GetUserByEmail(email)
    return append(users, u), err
  }
  u, err := ru.userRepository.GetUserByUsername(username)
  return append(users, u), err
}
