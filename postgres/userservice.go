package postgres

import (
	"database/sql"
	"fmt"
	"github.com/godwhoa/random-shit/botup.me/botup"
	_ "github.com/lib/pq"
	"log"
)

type UserService struct {
	DB *sql.DB
}

var create_stmt = "INSERT INTO USERS (UID,EMAIL,USERNAME,PASS) VALUES($1,$2,$3,$4)"

func (u *UserService) CreateUser(user User) error {
	err := u.DB.QueryRow(create_stmt, user.UID, user.Email, user.User, user.Pass)
	if err != nil {
		return botup.UserAlreadyExists
	}
	return nil
}

var get_stmt = "SELECT UID,EMAIL,USERNAME,PASS WHERE EMAIL = $1"

func (u *UserService) GetUser(email string) (User, error) {
	user := User{}
	err := u.DB.QueryRow(get_stmt, email).Scan(&user.UID, &user.Email, &user.Email, &user.Pass)
	if err == sql.ErrNoRows {
		return User{}, botup.UserDoesNotExist
	}
	return user, nil
}

var delete_user_stmt = "DELETE FROM USERS WHERE UID = $1"

func (u *UserService) DeleteUser(uid string) error {
	err := u.DB.QueryRow(delete_user_stmt, uid)
	if err != nil {
		return botup.UserDoesNotExist
	}
	return nil
}
