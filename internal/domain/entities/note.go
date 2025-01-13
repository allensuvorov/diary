package entities

import "time"

type Note struct {
	ID     int
	UserID int
	Time   time.Time
	Text   string
}
