package postgres

import (
	"database/sql"
	"github.com/godwhoa/random-shit/botup.me/botup"
	_ "github.com/lib/pq"
	"log"
)

type UserService struct {
	DB *sql.DB
}

var create_stmt = "INSERT INTO USERS (UID,EMAIL,USERNAME,PASS) VALUES($1,$2,$3,$4)"

func (u UserService) CreateUser(user botup.User) error {
	stmt, err := u.DB.Prepare(create_stmt)
	if err != nil {
		log.Println(err)
		return botup.UserAlreadyExists
	}
	_, err = stmt.Exec(user.UID, user.Email, user.User, user.Pass)
	return nil
}

var get_stmt = "SELECT UID,EMAIL,USERNAME,PASS WHERE EMAIL = $1"

func (u UserService) GetUser(email string) (botup.User, error) {
	user := botup.User{}
	err := u.DB.QueryRow(get_stmt, email).Scan(&user.UID, &user.Email, &user.Email, &user.Pass)
	if err == sql.ErrNoRows {
		return botup.User{}, botup.UserDoesNotExist
	}
	return user, nil
}

var delete_user_stmt = "DELETE FROM USERS WHERE UID = $1"

func (u UserService) DeleteUser(uid string) error {
	err := u.DB.QueryRow(delete_user_stmt, uid)
	if err != nil {
		return botup.UserDoesNotExist
	}
	return nil
}
