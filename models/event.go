package models

import "time"

type Event struct {
	Id          int
	Name        string
	Description string
	Location    string
	DateTime    time.Time
}

var events []Event

func (e *Event) Save() {
	events = append(events, *e)
}

func GetAllEvents() []Event {
	return events
}
