package entities

type User struct {
  Id       string
  Username string
  Password string
  Email    string
}

func NewUser(id, username, password, email string) User {
  return User{
    Id: id,
    Username: username,
    Password: password,
    Email: email,
  }
}
