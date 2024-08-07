package models

import (
	"database/sql"
	"rest-api/db"
	"time"
)

type Event struct {
	Id          int
	Name        string `binding:"required"`
	Description string `binding:"required"`
	Location    string `binding:"required"`
	DateTime    time.Time
	UserId      int
}

func (event Event) Save() error {
	insertQuery := `INSERT INTO events(name, description,location, dateTime,UserId) VALUES (?, ?, ?,?,?)`

	stmt, err := db.DB.Prepare(insertQuery)
	if err != nil {
		return err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {

		}
	}(stmt)
	result, err := stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.UserId)
	if err != nil {
		return err
	}
	_, err = result.LastInsertId()

	return err

}

func GetAllEvents() ([]Event, error) {
	selectQuery := `SELECT id,name,description,location,dateTime,UserId FROM events`

	rows, err := db.DB.Query(selectQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.Id, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}

func GetEventByID(id int64) (*Event, error) {
	var event Event

	selectEventQuery := `SELECT * FROM events WHERE id = ?`

	query, err := db.DB.Query(selectEventQuery, id)
	if err != nil {
		return nil, err
	}

	defer query.Close()

	query.Next()
	err = query.Scan(&event.Id, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)
	if err != nil {
		return nil, err
	}
	return &event, nil
}

func UpdateEventByID(id int64, event *Event) error {
	updateQuery := `
		UPDATE events
		SET name = ?, description = ?, location = ?, dateTime = ?
		WHERE id = ?
`
	stmt, err := db.DB.Prepare(updateQuery)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, id)
	if err != nil {
		return err
	}
	return err
}

func (event Event) Delete() error {
	deleteQuery := `DELETE FROM events WHERE id = ?`

	stmt, err := db.DB.Prepare(deleteQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(event.Id)
	if err != nil {
		return err
	}
	return err
}

func (event Event) Register(userId int) error {
	insertQuery := `INSERT INTO registrations(userId,eventId) VALUES(?,?)`

	_, err := db.DB.Exec(insertQuery, userId, event.Id)
	if err != nil {
		return err
	}
	return nil

}
