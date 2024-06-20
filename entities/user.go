package entities

type User struct {
  Id       string `gorm:"primaryKey"`
  Username string `gorm:"index:idx_username,unique"`
  Password string
  Email    string `gorm:"index:idx_user_email,unique"`
}

func NewUser(id, username, password, email string) User {
  return User{
    Id: id,
    Username: username,
    Password: password,
    Email: email,
  }
}
