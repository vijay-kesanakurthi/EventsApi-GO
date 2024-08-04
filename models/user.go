package models

import (
	"rest-api/db"
	"rest-api/util"
)

type User struct {
	Id       int
	Email    string
	Password string
}

func (user User) Save() error {
	insertUserQuery := `INSERT INTO users (email, password) VALUES (?, ?)`

	stmt, err := db.DB.Prepare(insertUserQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	hashedPassword, err := util.HashPassword(user.Password)
	_, err = stmt.Exec(user.Email, hashedPassword)
	if err != nil {
		return err
	}
	return err
}
