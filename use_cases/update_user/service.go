package updateuser

import (
  "fmt"
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
  if userRecovered.Id == "" {
    err = fmt.Errorf(uu.errorMsgs.USER_NOT_FOUND)
    return
  }
  if email == "" && password == "" {
    err = fmt.Errorf(uu.errorMsgs.USER_EMAIL_OR_PASSWORD_REQUIRED)
    return
  }
  if password != "" {
    userRecovered.Password = password
  }
  if email != "" {
    userRecovered.Email = email
  }
  u, err = uu.userRepository.UpdateUser(userRecovered)
  return
}
