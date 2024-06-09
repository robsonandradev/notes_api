package createuser

import (
  "log"
  "fmt"
	"github.com/robsonandradev/notes_api/config"
  repos "github.com/robsonandradev/notes_api/repositories"
  e "github.com/robsonandradev/notes_api/entities"
)

type CreateUserSvc struct {
	errorMsgs      config.ErrorMessages
  userRepository repos.IUserRepository
}

func NewCreateUserService(ur repos.IUserRepository) CreateUserSvc {
  return CreateUserSvc{
    userRepository: ur,
		errorMsgs:      *config.NewErrorMessages(),
  }
}

func (cr CreateUserSvc) Run(username, password, email string) (user e.User, err error) {
  log.Println("Registering new user:", username)
  if username == "" {
    return e.User{}, fmt.Errorf(cr.errorMsgs.USER_REQUIRED)
  }
  if password == "" {
    return e.User{}, fmt.Errorf(cr.errorMsgs.USER_PASSWORD_REQUIRED)
  }
  if email == "" {
    return e.User{}, fmt.Errorf(cr.errorMsgs.USER_EMAIL_REQUIRED)
  }
  user, err = cr.userRepository.CreateUser(username, password, email)
  return
}
