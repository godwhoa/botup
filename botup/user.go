package botup

import "errors"

type User struct {
	ID    int
	UID   string
	Email string
	User  string
	Pass  string
}

var UserAlreadyExists = errors.New("User already exists")
var UserDoesNotExist = errors.New("User does not exist")

type UserService interface {
	CreateUser(user User) error
	GetUser(email string) (User, error)
	DeleteUser(uid string) error
}
