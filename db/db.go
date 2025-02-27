package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./data.db")
	if err != nil {
		panic("failed to connect database")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	users := `CREATE TABLE IF NOT EXISTS users(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL
)`

	events := `CREATE TABLE IF NOT EXISTS events(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		userId INTEGER,
        FOREIGN KEY(userId) REFERENCES users(id)
	)`

	registrations := `CREATE TABLE IF NOT EXISTS registrations(
id INTEGER PRIMARY KEY AUTOINCREMENT,
userId INTEGER,
eventId INTEGER,
FOREIGN KEY(userId) REFERENCES users(id)
FOREIGN KEY(eventId) REFERENCES events(id)
)`
	_, err := DB.Exec(events)
	if err != nil {
		panic("failed to create events table: " + err.Error())
	}
	_, err = DB.Exec(users)
	if err != nil {
		panic("failed to create users table " + err.Error())
	}
	_, err = DB.Exec(registrations)
	if err != nil {
		panic("failed to create users table " + err.Error())
	}
}
