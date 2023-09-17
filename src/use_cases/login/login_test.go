package login

import (
  "testing"
)

func TestLoginWrongUsername(t *testing.T) {
  t.Run("when the username doen't exist then raise a error", func (t *testing.T) {
    login := New()
    want := "user not found!"
    _, err := login.doLogin("abc", "123")
    if err == nil || err.Error() != want{
      t.Errorf("expect %s error and got %s error", want, err)
    }
  })
}

func TestLoginWrongPassword(t *testing.T) {
  t.Run("when the password is wrong it should raise an error", func(t *testing.T) {
    login := New()
    want := "wrong password!"
    _, err := login.doLogin("john", "123")
    if err == nil || err.Error() != want{
      t.Errorf("expect %s error and got %s error", want, err)
    }
  })
}
