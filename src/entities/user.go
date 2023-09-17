package entities

type User struct {
  id       string
  username string
  password string
  email    string
}

func (u User) New() User {
  return User{}
}
