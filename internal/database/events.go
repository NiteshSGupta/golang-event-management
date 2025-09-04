package database

import "database/sql"

type EventModel struct {
	DB *sql.DB
}

type Events struct {
	Id          int    `json:"id"`
	OwnerId     string `json:"owner_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Date        string `json:"date"`
	Location    string `json:"location"`
}
