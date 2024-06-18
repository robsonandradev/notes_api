package readuser

import (
  "os"
  "testing"
	"github.com/robsonandradev/notes_api/config"
  e "github.com/robsonandradev/notes_api/entities"
)

var (
  readUserSvc ReadUserSvc
  errorMsgs config.ErrorMessages
)

func TestMain(m *testing.M) {
  mock := UserRepositoryMock{}
  errorMsgs = *config.NewErrorMessages()
  readUserSvc = *NewReadUserSvc(&mock)
  os.Exit(m.Run())
}

func TestReadUserSuccessfuly(t *testing.T) {
  t.Run("when get user by id and user exists", func (t *testing.T) {
    userId := "1"
    mockedUsers := mockUsers()
    want := mockedUsers[0]
    users, err := readUserSvc.Run("", "", userId)
    if err != nil { panic(err) }
    if users[0] != want {
      t.Errorf("expect %s and got %s", want, users[0])
    }
  })
  t.Run("when get user by username and user exists", func (t *testing.T) {
    username := "john"
    mockedUsers := mockUsers()
    want := mockedUsers[0]
    users, err := readUserSvc.Run(username, "", "")
    if err != nil { panic(err) }
    if users[0] != want {
      t.Errorf("expect %s and got %s", want, users[0])
    }
  })
  t.Run("when get user by email and user exists", func (t *testing.T) {
    email := "john@gmail.com"
    mockedUsers := mockUsers()
    want := mockedUsers[0]
    users, err := readUserSvc.Run("", email, "")
    if err != nil { panic(err) }
    if users[0] != want {
      t.Errorf("expect %s and got %s", want, users[0])
    }
  })
  t.Run("when get all users", func (t *testing.T) {
    want := mockUsers()
    users, err := readUserSvc.Run("", "", "")
    if err != nil { panic(err) }
    for k, v := range users {
      if v != want[k] {
        t.Errorf("expect %s and got %s", want[k], v)
      }
    }
  })
}

func TestReadUserNotFound(t *testing.T) {
  t.Run("when get user by id and user doesnt exists", func(t *testing.T) {
    userId := "55"
    want := append([]e.User{}, e.User{})
    users, err := readUserSvc.Run("", "", userId)
    if err != nil { panic(err) }
    if len(want) != len(users) {
      t.Errorf("expect %s and got %s", want, users)
    }
  })
  t.Run("when get user by username and user doesnt exists", func(t *testing.T) {
    username := "user@not.found"
    want := append([]e.User{}, e.User{})
    users, err := readUserSvc.Run(username, "", "")
    if err != nil { panic(err) }
    if len(want) != len(users) {
      t.Errorf("expect %s and got %s", want, users)
    }
  })
  t.Run("when get user by email and user doesnt exists", func(t *testing.T) {
    email := "user@not.found"
    want := append([]e.User{}, e.User{})
    users, err := readUserSvc.Run("", email, "")
    if err != nil { panic(err) }
    if len(want) != len(users) {
      t.Errorf("expect %s and got %s", want, users)
    }
  })
}
// starting mocks
type UserRepositoryMock struct {}

func (ur *UserRepositoryMock) GetUserByUsername(username string) (u e.User, e error) {
  users := mockUsers()
  for _, v := range users {
    if v.Username == username {
      u = v
      return
    }
  }
  return
}

func (ur *UserRepositoryMock) GetUserByEmail(email string) (u e.User, err error) {
  users := mockUsers()
  for _, v := range users {
    if v.Email == email {
      u = v
      return
    }
  }
  return
}

func (ur *UserRepositoryMock) GetUserById(id string) (u e.User, err error) {
  users := mockUsers()
  for _, v := range users {
    if v.Id == id {
      u = v
      return
    }
  }
  return
}

func (ur *UserRepositoryMock) GetAll() (users []e.User, err error) {
  return mockUsers(), err
}

func (ur *UserRepositoryMock) GetUserByUsernameAndPassword(username, password string) (u e.User, e error) { return }

func (ur *UserRepositoryMock) CreateUser(username, password, email string) (u e.User, e error) { return }

func (ur * UserRepositoryMock) UpdateUser(newUser e.User) (u e.User, err error) { return }

func mockUsers() []e.User {
  u1 := e.NewUser("1", "john", "john", "john@gmail.com")
  u2 := e.NewUser("2", "jone", "jone", "jone@gmail.com")
  u3 := e.NewUser("3", "joel", "joel", "joel@gmail.com")
  return []e.User{u1, u2, u3}
}
// ending mocks
