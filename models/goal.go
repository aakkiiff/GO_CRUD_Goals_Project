package models

import "time"

type Goal struct {
	ID          int64
	Name        string
	Description string
	DateTime    time.Time
}
