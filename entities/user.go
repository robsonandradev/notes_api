package entities

type User struct {
  id       string
  username string
  password string
  email    string
}

func (u User) New(id, username, password, email string) User {
  return User{
    id: id,
    username: username,
    password: password,
    email: email,
  }
}
