package postgres

import (
	"time"

	"gorm.io/gorm"
)

type BookEntity struct {
	gorm.Model
	ID          string `gorm:"primaryKey"`
	Title       string
	Author      string
	Topic       string
	Year        int
	Popularity  int
	IsLoaned    bool
	LoanedTo    string
	LoanedUntil time.Time
}
