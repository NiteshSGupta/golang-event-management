package database

import "database/sql"

type AttendeeModel struct {
	DB *sql.DB
}

// data in response
type Attendees struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"-"`
}
