package models

import (
	"database/sql"
	"rest-api/db"
	"time"
)

type Event struct {
	Id          int
	Name        string
	Description string
	Location    string
	DateTime    time.Time
}

var events []Event

func (e *Event) Save() error {
	insertQuery := `INSERT INTO events(name, description,location, dateTime) VALUES (?, ?, ?,?)`

	stmt, err := db.DB.Prepare(insertQuery)
	if err != nil {
		return err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {

		}
	}(stmt)
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime)
	if err != nil {
		return err
	}
	_, err = result.LastInsertId()

	return err

}

func GetAllEvents() ([]Event, error) {
	selectQuery := `SELECT * FROM events`

	rows, err := db.DB.Query(selectQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.Id, &event.Name, &event.Description, &event.Location, &event.DateTime)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}
