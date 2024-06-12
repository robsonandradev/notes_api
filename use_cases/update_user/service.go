package updateuser

import (
	"github.com/robsonandradev/notes_api/config"
  e "github.com/robsonandradev/notes_api/entities"
  repos "github.com/robsonandradev/notes_api/repositories"
)

type UpdateUserSvc struct {
	errorMsgs      config.ErrorMessages
  userRepository repos.IUserRepository
}

func NewUpdateUserSvc(repo repos.IUserRepository) *UpdateUserSvc {
  return &UpdateUserSvc{
    errorMsgs: *config.NewErrorMessages(),
    userRepository: repo,
  }
}

func (uu UpdateUserSvc) Run(id, password, email string) (u e.User, err error) {
  userRecovered, err := uu.userRepository.GetUserById(id)
  if err != nil { return }
  if password == "" {
    password = userRecovered.Password
  }
  if email == "" {
    email = userRecovered.Email
  }
  u, err = uu.userRepository.UpdateUser(id, password, email)
  return
}
