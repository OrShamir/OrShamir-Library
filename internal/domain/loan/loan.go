package loan

import "time"

type Loan struct {
	ID         string
	BookID     string
	UserID     string
	LoanedAt   time.Time
	DueDate    time.Time
	ReturnedAt *time.Time // Null if not returned
}
