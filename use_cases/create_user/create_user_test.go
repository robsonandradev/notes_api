package createuser

import (
  "os"
  "testing"
	"github.com/robsonandradev/notes_api/config"
  e "github.com/robsonandradev/notes_api/entities"
)

var (
  createUserSvc CreateUserSvc
  errorMsgs config.ErrorMessages
)

func TestMain(m *testing.M) {
  mock := UserRepositoryMock{}
  errorMsgs = *config.NewErrorMessages()
  createUserSvc = NewCreateUserService(&mock)
  os.Exit(m.Run())
}

func TestSuccessfulUser(t *testing.T) {
  t.Run("when add new user with all correct data", func (t *testing.T) {
    want := mockUsers()[0]
    user, err := createUserSvc.Run(want.Username, want.Password, want.Email)
    if err != nil {
      t.Errorf("expect %s error and got %s error", want, err)
    }
    if user.Id != want.Id {
      t.Errorf("expect %s error and got %s error", want.Id, user.Id)
    }
  })
}

func TestLoginWrongUsername(t *testing.T) {
  t.Run("when add new user without username", func (t *testing.T) {
    want := errorMsgs.USER_REQUIRED
    _, err := createUserSvc.Run("", "john!123", "john.wick@gmail.com")
    if err == nil || err.Error() != want{
      t.Errorf("expect %s error and got %s error", want, err)
    }
  })
  t.Run("when add new user without password", func (t *testing.T) {
    want := errorMsgs.USER_PASSWORD_REQUIRED
    _, err := createUserSvc.Run("john.wick", "", "john.wick@gmail.com")
    if err == nil || err.Error() != want{
      t.Errorf("expect %s error and got %s error", want, err)
    }
  })
  t.Run("when add new user without email", func (t *testing.T) {
    want := errorMsgs.USER_EMAIL_REQUIRED
    _, err := createUserSvc.Run("john.wick", "john!123", "")
    if err == nil || err.Error() != want{
      t.Errorf("expect %s error and got %s error", want, err)
    }
  })
}

// starting mocks
type UserRepositoryMock struct {}

func (ur *UserRepositoryMock) GetUserByUsername(username string) (u e.User, e error) { return }

func (ur *UserRepositoryMock) GetUserByEmail(email string) (u e.User, err error) { return }

func (ur *UserRepositoryMock) GetUserById(id string) (u e.User, err error) { return }

func (ur *UserRepositoryMock) GetAll() (users []e.User, err error) {return}

func (ur *UserRepositoryMock) GetUserByUsernameAndPassword(username, password string) (u e.User, e error) { return }

func (ur *UserRepositoryMock) CreateUser(username, password, email string) (u e.User, e error) {
  return mockUsers()[0], nil
}

func (ur * UserRepositoryMock) UpdateUser(id, password, email string) (u e.User, err error) { return }

func mockUsers() []e.User {
  u1 := e.NewUser("1", "john", "john", "john@gmail.com")
  u2 := e.NewUser("2", "jone", "jone", "jone@gmail.com")
  u3 := e.NewUser("3", "joel", "joel", "joel@gmail.com")
  return []e.User{u1, u2, u3}
}
// ending mocks
