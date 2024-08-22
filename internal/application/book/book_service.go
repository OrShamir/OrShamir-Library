package book

import (
	"Or/Library/internal/domain/book"
	"context"
	"errors"
	"time"
)

type BookService struct {
	bookRepository book.BookRepository
}

func NewBookService(bookRepository book.BookRepository) *BookService {
	return &BookService{bookRepository}
}

func (s *BookService) CreateBook(ctx context.Context, book *book.Book) error {
	if book.Title == "" || book.Author == "" {
		return errors.New("title and author are required")
	}
	if book.Popularity < 1 || book.Popularity > 5 {
		return errors.New("popularity must be between 1 and 5")
	}

	return s.bookRepository.Create(ctx, book)
}

func (s *BookService) GetBook(ctx context.Context, id string) (*book.Book, error) {
	return s.bookRepository.GetByID(ctx, id)
}

func (s *BookService) UpdateBook(ctx context.Context, book *book.Book) error {
	if book.IsLoaned {
		return errors.New("cannot update a loaned book")
	}
	// ... other validations (e.g., title, author, popularity)

	return s.bookRepository.Update(ctx, book)
}

func (s *BookService) DeleteBook(ctx context.Context, id string) error {
	book, err := s.bookRepository.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if book.IsLoaned {
		return errors.New("cannot delete a loaned book")
	}

	return s.bookRepository.Delete(ctx, id)
}

func (s *BookService) SearchBooks(ctx context.Context, query string) ([]*book.Book, error) {
	return s.bookRepository.Search(ctx, query)
}

func (s *BookService) LoanBook(ctx context.Context, bookID, userID string) error {
	// 1. Get the book
	book, err := s.bookRepository.GetByID(ctx, bookID)
	if err != nil {
		return err
	}

	// 2. Check if the book is available
	if book.IsLoaned {
		return errors.New("book is already loaned out")
	}

	// 3. Update book loan details
	book.IsLoaned = true
	book.LoanedTo = userID
	book.LoanedUntil = time.Now().AddDate(0, 0, book.LoanDuration())

	// 4. Save changes
	return s.bookRepository.Update(ctx, book)
}

func (s *BookService) ReturnBook(ctx context.Context, bookID string) error {
	// 1. Get the book
	book, err := s.bookRepository.GetByID(ctx, bookID)
	if err != nil {
		return err
	}

	// 2. Reset book loan details
	book.IsLoaned = false
	book.LoanedTo = ""
	book.LoanedUntil = time.Time{}

	// 3. Save changes
	return s.bookRepository.Update(ctx, book)
}
