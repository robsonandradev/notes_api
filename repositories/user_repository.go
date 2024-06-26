package repositories

import (
  "fmt"
  "github.com/google/uuid"
  e "github.com/robsonandradev/notes_api/entities"
  "github.com/robsonandradev/notes_api/config"
)

type IUserRepository interface {
  GetUserByUsernameAndPassword(username, password string) (e.User, error)
  GetUserByUsername(username string) (e.User, error)
  GetUserByEmail(email string) (e.User, error)
  GetUserById(id string) (e.User, error)
  GetAll() ([]e.User, error)
  CreateUser(username, password, email string) (e.User, error)
  UpdateUser(newUser e.User) (e.User, error)
}

type UserRepository struct {
  db DBConnection
}

func NewUserRepository(connector string) (*UserRepository, error) {
  consts := config.NewConstants()
  if connector == consts.POSTGRES {
    ur := &UserRepository{ db: &PostgresCon{} }
    return ur, nil
  }
  return &UserRepository{}, fmt.Errorf("Unknow database connector")
}

func (ur *UserRepository) GetUserByUsername(username string) (e.User, error) {
  db, err := ur.db.Connect()
  if err != nil {
    return e.User{}, err
  }
  defer ur.db.Close(db)
  user := e.User{}
  r := db.First(&user, "username = ?", username)
  return user, r.Error
}

func (ur *UserRepository) GetUserByUsernameAndPassword(username, password string) (e.User, error) {
  db, err := ur.db.Connect()
  if err != nil {
    return e.User{}, err
  }
  defer ur.db.Close(db)
  user := e.User{}
  r := db.First(&user, "username = ? and password = ?", username, password)
  return user, r.Error
}

func (ur *UserRepository) GetUserByEmail(email string) (u e.User, err error) {
	db, err := ur.db.Connect()
	if err != nil {
		return
	}
	defer ur.db.Close(db)
	r := db.Where("email = ?", email).Find(&u)
	return u, r.Error
}

func (ur *UserRepository) GetUserById(id string) (u e.User, err error) {
	db, err := ur.db.Connect()
	if err != nil {
		return
	}
	defer ur.db.Close(db)
	r := db.Where("id = ?", id).Find(&u)
	return u, r.Error
}

func (ur *UserRepository) GetAll() (users []e.User, err error) {
	db, err := ur.db.Connect()
	if err != nil {
		return
	}
	defer ur.db.Close(db)
	r := db.Find(&users)
	return users, r.Error
}

func (ur *UserRepository) CreateUser(username, password, email string) (u e.User, err error) {
  db, err := ur.db.Connect()
  if err != nil {
    return 
  }
  defer ur.db.Close(db)
  u.Id       = uuid.NewString()
  u.Email    = email
  u.Username = username
  u.Password = password
  r := db.Create(&u)
  return u, r.Error
}

func (ur * UserRepository) UpdateUser(newUser e.User) (u e.User, err error) {
  db, err := ur.db.Connect()
  if err != nil {
    return
  }
  defer ur.db.Close(db)
  r := db.Model(&newUser).Updates(newUser)
  return newUser, r.Error
}
