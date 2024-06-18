package updateuser

import (
  "fmt"
  "os"
  "testing"
	"github.com/robsonandradev/notes_api/config"
  e "github.com/robsonandradev/notes_api/entities"
)

var (
  updateuserSvc *UpdateUserSvc
  errorMsgs config.ErrorMessages
)

func TestMain(m *testing.M) {
  mock := &UserRepositoryMock{}
  errorMsgs = *config.NewErrorMessages()
  updateuserSvc = NewUpdateUserSvc(mock)
  os.Exit(m.Run())
}

func TestUpdateUserSuccessfuly(t *testing.T) {
  t.Run("when user change password with correct data", func (t *testing.T) {
    newPassword := "mudar123"
    want := mockUsers()[0]
    user, err := updateuserSvc.Run(want.Id, newPassword, want.Email)
    if err != nil { panic(err) }
    if user.Id != want.Id {
      t.Errorf("expected %s and got %s", want, user)
    }
    if user.Password == want.Password {
      t.Errorf("expected %s and got %s", newPassword, user.Password)
    }
  })
  t.Run("when user change password without send email", func (t *testing.T) {
    newPassword := "mudar123"
    want := mockUsers()[0]
    user, err := updateuserSvc.Run(want.Id, newPassword, "")
    if err != nil { panic(err) }
    if user.Id != want.Id {
      t.Errorf("expected %s and got %s", want, user)
    }
    if user.Password == want.Password {
      t.Errorf("expected %s and got %s", newPassword, user.Password)
    }
    if user.Email != want.Email {
      t.Errorf("expected %s and got %s", want.Email, user.Email)
    }
  })
  t.Run("when user change email with correct data", func (t *testing.T) {
    newEmail := "new@mail.com"
    want := mockUsers()[0]
    user, err := updateuserSvc.Run(want.Id, want.Password, newEmail)
    if err != nil { panic(err) }
    if user.Id != want.Id {
      t.Errorf("expected %s and got %s", want, user)
    }
    if user.Email == want.Email {
      t.Errorf("expected %s and got %s", newEmail, user.Email)
    }
  })
  t.Run("when user change email without send password", func (t *testing.T) {
    newEmail := "new@mail.com"
    want := mockUsers()[0]
    user, err := updateuserSvc.Run(want.Id, "", newEmail)
    if err != nil { panic(err) }
    if user.Id != want.Id {
      t.Errorf("expected %s and got %s", want, user)
    }
    if user.Email == want.Email {
      t.Errorf("expected %s and got %s", newEmail, user.Email)
    }
    if user.Password != want.Password {
      t.Errorf("expected %s and got %s", want.Password, user.Password)
    }
  })
}

func TestUpdateUserFail(t *testing.T) {
  t.Run("when update user that doesnt exists", func (t *testing.T) {
    randomId := "123654"
    want := fmt.Errorf(errorMsgs.USER_NOT_FOUND)
    _, err := updateuserSvc.Run(randomId, "any", "any")
    if err.Error() != want.Error() {
      t.Errorf("expected %s exception and got %s", want, err)
    }
  })
  t.Run("when update user without send email and password", func (t *testing.T) {
    mockedUser := mockUsers()[0]
    want := fmt.Errorf(errorMsgs.USER_EMAIL_OR_PASSWORD_REQUIRED)
    _, err := updateuserSvc.Run(mockedUser.Id, "", "")
    if err == nil || err.Error() != want.Error() {
      t.Errorf("expected %s exception and got %s", want, err)
    }
  })
}

// starting mocks
type UserRepositoryMock struct {}

func (ur *UserRepositoryMock) GetUserByUsername(username string) (u e.User, e error) { return }
func (ur *UserRepositoryMock) GetUserByEmail(email string) (u e.User, err error) { return }
func (ur *UserRepositoryMock) GetUserById(id string) (u e.User, err error) {
  users := mockUsers()
  for _, u = range users {
    if u.Id == id {
      return
    }
  }
  u = e.User{}
  return
}
func (ur *UserRepositoryMock) GetAll() (users []e.User, err error) { return }
func (ur *UserRepositoryMock) GetUserByUsernameAndPassword(username, password string) (u e.User, e error) { return }
func (ur *UserRepositoryMock) CreateUser(username, password, email string) (u e.User, e error) { return }
func (ur * UserRepositoryMock) UpdateUser(newUser e.User) (u e.User, err error) {
  u = mockUsers()[0]
  u.Password = newUser.Password
  u.Email = newUser.Email
  return
}

func mockUsers() []e.User {
  u1 := e.NewUser("1", "john", "john", "john@gmail.com")
  u2 := e.NewUser("2", "jone", "jone", "jone@gmail.com")
  u3 := e.NewUser("3", "joel", "joel", "joel@gmail.com")
  return []e.User{u1, u2, u3}
}
// ending mocks
