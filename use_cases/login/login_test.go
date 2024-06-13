package login

import (
  "os"
  "fmt"
  "testing"
  e "github.com/robsonandradev/notes_api/entities"
)

var login loginSvc

func TestMain(m *testing.M) {
  mock := UserRepositoryMock{}
  login = New(&mock)
  os.Exit(m.Run())
}

func TestLoginWrongUsername(t *testing.T) {
  t.Run("when the username doen't exist then raise a error", func (t *testing.T) {
    want := "user not found!"
    _, err := login.doLogin("abc", "123")
    if err == nil || err.Error() != want{
      t.Errorf("expect %s error and got %s error", want, err)
    }
  })
}

func TestLoginWrongPassword(t *testing.T) {
  t.Run("when the password is wrong it should raise an error", func(t *testing.T) {
    want := "wrong password!"
    _, err := login.doLogin("john", "123")
    if err == nil || err.Error() != want{
      t.Errorf("expect %s error and got %s error", want, err)
    }
  })
}

func TestSuccessfulLogin(t *testing.T) {
  t.Run("when username and password is correct return the user", func(t *testing.T) {
    want := e.NewUser("1", "john", "john", "john@gmail.com")
    have, _ := login.doLogin("john", "john")
    if want != have {
      t.Errorf("expect %s and have %s", want, have)
    }
  })
}

// starting mocks
type UserRepositoryMock struct {}

func (ur *UserRepositoryMock) GetUserByUsername(username string) (e.User, error) {
  users := mockUsers()
  for _, user := range users {
    if user.Username == username {
      return user, nil
    }
  }
  return e.User{}, fmt.Errorf("user not found!")
}

func (ur *UserRepositoryMock) GetUserByEmail(email string) (u e.User, err error) { return }

func (ur *UserRepositoryMock) GetUserById(id string) (u e.User, err error) { return }

func (ur *UserRepositoryMock) GetAll() (users []e.User, err error) {return}

func (ur *UserRepositoryMock) GetUserByUsernameAndPassword(username, password string) (e.User, error) {
  users := mockUsers()
  for _, user := range users {
    if user.Username == username && user.Password == password {
        return user, nil
    }
  }
  return e.User{}, fmt.Errorf("wrong password!")
}

func (ur *UserRepositoryMock) CreateUser(username, password, email string) (u e.User, e error) { return }
func (ur * UserRepositoryMock) UpdateUser(id, password, email string) (u e.User, err error) { return }

func mockUsers() []e.User {
  u1 := e.NewUser("1", "john", "john", "john@gmail.com")
  u2 := e.NewUser("2", "jone", "jone", "jone@gmail.com")
  u3 := e.NewUser("3", "joel", "joel", "joel@gmail.com")
  return []e.User{u1, u2, u3}
}
// ending mocks
