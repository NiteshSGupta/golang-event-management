package database

import "database/sql"

type AttendeeModel struct {
	DB *sql.DB
}

// data in response
type Attendees struct {
	Id      int `json:"id"`
	UserId  int `json:"user_id"`
	EventId int `json:"event_id"`
}
