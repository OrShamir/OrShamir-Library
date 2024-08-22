package book

import "time"

type Book struct {
	ID          string
	Title       string
	Author      string
	Topic       string
	Year        int
	Popularity  int // 1-5 stars
	IsLoaned    bool
	LoanedTo    string    // User ID if loaned
	LoanedUntil time.Time // Date if loaned
}

func (b *Book) LoanDuration() int {
	switch b.Popularity {
	case 1, 2, 3:
		return 7 // 1 week
	case 4:
		return 3 // 3 days
	case 5:
		return 2 // 2 days
	default:
		return 7 // Default to 1 week
	}
}
