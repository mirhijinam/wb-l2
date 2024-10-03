package models

import "time"

type Event struct {
	ID   int
	Name string
	Data string
	Date time.Time
}
