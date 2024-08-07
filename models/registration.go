package models

import "rest-api/db"

type Registration struct {
	Id      int `json:"id"`
	UserId  int `json:"user_id"`
	EventId int `json:"event_id"`
}

type RegistrationsData struct {
	Id               int    `json:"id"`
	UserId           int    `json:"user_id"`
	EventId          int    `json:"event_id"`
	EventName        string `json:"event_name"`
	EventDescription string `json:"event_description"`
}

func (reg Registration) Delete() error {
	query := "DELETE FROM registrations WHERE userId=? AND eventId=?"

	_, err := db.DB.Exec(query, reg.UserId, reg.EventId)
	if err != nil {
		return err
	}
	return err
}

func FindRegistrationById(id int64) (Registration, error) {
	var reg Registration

	query := `SELECT * FROM registrations WHERE id=?`

	row := db.DB.QueryRow(query, id)
	err := row.Scan(&reg.Id, &reg.UserId, &reg.EventId)
	if err != nil {
		return reg, err
	}
	return reg, nil
}

func FindAllRegistrations() ([]RegistrationsData, error) {
	var regs []RegistrationsData
	query := `SELECT registrations.id,registrations.userId,events.id,events.name,events.description FROM registrations INNER JOIN events ON registrations.eventId=events.id `
	rows, err := db.DB.Query(query)
	if err != nil {
		return regs, err
	}
	defer rows.Close()
	for rows.Next() {
		var reg RegistrationsData
		err := rows.Scan(&reg.Id, &reg.UserId, &reg.EventId, &reg.EventName, &reg.EventDescription)
		if err != nil {
			return nil, err
		}
		regs = append(regs, reg)
	}
	return regs, nil
}

func FindRegistrationsByUserId(userId int) ([]Registration, error) {
	var regs []Registration
	query := `SELECT * FROM registrations WHERE userId=? `
	rows, err := db.DB.Query(query, userId)
	if err != nil {
		return regs, err
	}
	defer rows.Close()
	for rows.Next() {
		var reg Registration
		err := rows.Scan(&reg.Id, &reg.UserId, &reg.EventId)
		if err != nil {
			return regs, err
		}
	}

	return regs, nil

}
