package loan

import (
	"context"
)

type LoanRepository interface {
	Create(ctx context.Context, loan *Loan) error
	GetByID(ctx context.Context, id string) (*Loan, error)
	Update(ctx context.Context, loan *Loan) error
	Delete(ctx context.Context, id string) error
	GetByUser(ctx context.Context, userID string) ([]Loan, error)
	GetByBook(ctx context.Context, bookID string) (*Loan, error)
}
