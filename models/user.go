package models

import (
	"rest-api/db"
	"rest-api/util"
)

type User struct {
	Id       int    `json:"id" `
	Email    string `binding:"required" json:"email"`
	Password string `binding:"required" json:"password"`
}

type UserModel struct {
	Email    string `binding:"required" json:"email"`
	Password string `binding:"required" json:"password"`
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

func Validate(user *User) error {
	findUserQuery := `SELECT * FROM users WHERE email = ?`

	var actualUser User

	rows, err := db.DB.Query(findUserQuery, user.Email)
	if err != nil {
		return err
	}
	defer rows.Close()
	rows.Next()
	err = rows.Scan(&actualUser.Id, &actualUser.Email, &actualUser.Password)
	if err != nil {
		return err
	}
	user.Id = actualUser.Id

	result := util.ComparePasswords(actualUser.Password, user.Password)
	if result == false {
		return err
	}
	return nil

}
