package loan

import (
	"Or/Library/internal/domain/book"
	"Or/Library/internal/domain/loan"
	"Or/Library/internal/domain/user"
	"context"
	"errors"
	"time"
)

type LoanService struct {
	loanRepository loan.LoanRepository
	bookRepository book.BookRepository
	userRepository user.UserRepository
}

func NewLoanService(loanRepository loan.LoanRepository, bookRepository book.BookRepository, userRepository user.UserRepository) *LoanService {
	return &LoanService{loanRepository, bookRepository, userRepository}
}

func (s *LoanService) CreateLoan(ctx context.Context, loan *loan.Loan) error {
	// 1. Get the book and user
	b, err := s.bookRepository.GetByID(ctx, loan.BookID)
	if err != nil {
		return err
	}
	u, err := s.userRepository.GetByID(ctx, loan.UserID)
	if err != nil {
		return err
	}

	// 2. Check if the book is available and the user has less than 5 loaned books
	if b.IsLoaned {
		return errors.New("book is already loaned out")
	}
	userLoans, err := s.loanRepository.GetByUser(ctx, u.ID)
	if err != nil {
		return err
	}
	if len(userLoans) >= 5 {
		return errors.New("user has reached the maximum loan limit")
	}

	// 3. Update book loan details
	b.IsLoaned = true
	b.LoanedTo = u.ID
	b.LoanedUntil = time.Now().AddDate(0, 0, b.LoanDuration())
	err = s.bookRepository.Update(ctx, b)
	if err != nil {
		return err
	}

	// 4. Create the loan record
	loan.LoanedAt = time.Now()
	loan.DueDate = b.LoanedUntil
	return s.loanRepository.Create(ctx, loan)
}

func (s *LoanService) GetLoan(ctx context.Context, id string) (*loan.Loan, error) {
	return s.loanRepository.GetByID(ctx, id)
}

func (s *LoanService) ReturnBook(ctx context.Context, loanID string) error {
	// 1. Get the loan
	l, err := s.loanRepository.GetByID(ctx, loanID)
	if err != nil {
		return err
	}

	// 2. Get the book
	b, err := s.bookRepository.GetByID(ctx, l.BookID)
	if err != nil {
		return err
	}

	// 3. Reset book loan details
	b.IsLoaned = false
	b.LoanedTo = ""
	b.LoanedUntil = time.Time{}
	err = s.bookRepository.Update(ctx, b)
	if err != nil {
		return err
	}

	// 4. Update the loan record
	now := time.Now()
	l.ReturnedAt = &now
	return s.loanRepository.Update(ctx, l)
}
