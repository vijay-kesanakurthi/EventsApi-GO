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

	createEventTable()
}

func createEventTable() {
	table := `CREATE TABLE IF NOT EXISTS events(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL
	)`
	_, err := DB.Exec(table)
	if err != nil {
		panic("failed to create events table")
	}
}
