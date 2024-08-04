package models

import "rest-api/db"

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

	_, err = stmt.Exec(user.Email, user.Password)
	if err != nil {
		return err
	}
	return err
}
