package loan

import (
	"time"

	"gorm.io/gorm"
)

type LoanEntity struct {
	gorm.Model
	ID         string `gorm:"primaryKey"`
	BookID     string
	UserID     string
	LoanedAt   time.Time
	DueDate    time.Time
	ReturnedAt *time.Time // Null if not returned
}
