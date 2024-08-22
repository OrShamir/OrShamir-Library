package book

import (
	"context"
	"testing"
	"time"

	"Or/Library/internal/domain/book"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockBookRepository struct {
	mock.Mock
}

func (m *mockBookRepository) Create(ctx context.Context, book *book.Book) error {
	args := m.Called(ctx, book)
	return args.Error(0)
}

func (m *mockBookRepository) GetByID(ctx context.Context, id string) (*book.Book, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*book.Book), args.Error(1)
}

func (m *mockBookRepository) Update(ctx context.Context, book *book.Book) error {
	args := m.Called(ctx, book)
	return args.Error(0)
}

func (m *mockBookRepository) Delete(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *mockBookRepository) Search(ctx context.Context, query string) ([]*book.Book, error) {
	args := m.Called(ctx, query)
	return args.Get(0).([]*book.Book), args.Error(1)
}

func TestCreateBook(t *testing.T) {
	mockRepo := new(mockBookRepository)
	service := NewBookService(mockRepo)

	t.Run("valid book", func(t *testing.T) {
		b := &book.Book{
			Title:      "Test Book",
			Author:     "Test Author",
			Popularity: 3,
		}
		mockRepo.On("Create", context.Background(), b).Return(nil)

		err := service.CreateBook(context.Background(), b)
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	// ... other test cases for invalid inputs and repository errors
}

func TestGetBook(t *testing.T) {
	mockRepo := new(mockBookRepository)
	service := NewBookService(mockRepo)

	t.Run("existing book", func(t *testing.T) {
		b := &book.Book{
			ID:    "123",
			Title: "Test Book",
		}
		mockRepo.On("GetByID", context.Background(), "123").Return(b, nil)

		result, err := service.GetBook(context.Background(), "123")
		assert.NoError(t, err)
		assert.Equal(t, b, result)
		mockRepo.AssertExpectations(t)
	})

	// ... other test cases for non-existent book and repository errors
}

func TestUpdateBook(t *testing.T) {
	mockRepo := new(mockBookRepository)
	service := NewBookService(mockRepo)

	t.Run("valid update", func(t *testing.T) {
		b := &book.Book{
			ID:         "123",
			Title:      "Test Book",
			IsLoaned:   false,
			Popularity: 3,
		}
		updatedBook := &book.Book{
			ID:         "123",
			Title:      "Updated Title",
			IsLoaned:   false,
			Popularity: 4,
		}
		mockRepo.On("GetByID", context.Background(), "123").Return(b, nil)
		mockRepo.On("Update", context.Background(), updatedBook).Return(nil)

		err := service.UpdateBook(context.Background(), updatedBook)
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	// ... other test cases for loaned book, invalid inputs, and repository errors
}

func TestDeleteBook(t *testing.T) {
	mockRepo := new(mockBookRepository)
	service := NewBookService(mockRepo)

	t.Run("valid delete", func(t *testing.T) {
		b := &book.Book{
			ID:       "123",
			IsLoaned: false,
		}
		mockRepo.On("GetByID", context.Background(), "123").Return(b, nil)
		mockRepo.On("Delete", context.Background(), "123").Return(nil)

		err := service.DeleteBook(context.Background(), "123")
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	// ... other test cases for loaned book, non-existent book, and repository errors
}

func TestSearchBooks(t *testing.T) {
	mockRepo := new(mockBookRepository)
	service := NewBookService(mockRepo)

	t.Run("search by title", func(t *testing.T) {
		books := []*book.Book{
			{ID: "1", Title: "Book 1"},
			{ID: "2", Title: "Another Book"},
		}
		mockRepo.On("Search", context.Background(), "Book").Return(books, nil)

		result, err := service.SearchBooks(context.Background(), "Book")
		assert.NoError(t, err)
		assert.Equal(t, books, result)
		mockRepo.AssertExpectations(t)
	})

	// ... other test cases for search by author, topic, and repository errors
}

func TestLoanBook(t *testing.T) {
	mockRepo := new(mockBookRepository)
	service := NewBookService(mockRepo)

	t.Run("successful loan", func(t *testing.T) {
		b := &book.Book{
			ID:         "123",
			IsLoaned:   false,
			Popularity: 3, // 1 week loan duration
		}
		mockRepo.On("GetByID", context.Background(), "123").Return(b, nil)
		mockRepo.On("Update", context.Background(), mock.AnythingOfType("*book.Book")).Return(nil).
			Run(func(args mock.Arguments) {
				updatedBook := args.Get(1).(*book.Book)
				assert.True(t, updatedBook.IsLoaned)
				assert.Equal(t, "user123", updatedBook.LoanedTo)
				assert.WithinDuration(t, time.Now().AddDate(0, 0, 7), updatedBook.LoanedUntil, time.Minute)
			})

		err := service.LoanBook(context.Background(), "123", "user123")
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	// ... other test cases for already loaned book, repository errors
}

func TestReturnBook(t *testing.T) {
	mockRepo := new(mockBookRepository)
	service := NewBookService(mockRepo)

	t.Run("successful return", func(t *testing.T) {
		b := &book.Book{
			ID:       "123",
			IsLoaned: true,
			LoanedTo: "user123",
		}
		mockRepo.On("GetByID", context.Background(), "123").Return(b, nil)
		mockRepo.On("Update", context.Background(), mock.AnythingOfType("*book.Book")).Return(nil).
			Run(func(args mock.Arguments) {
				updatedBook := args.Get(1).(*book.Book)
				assert.False(t, updatedBook.IsLoaned)
				assert.Empty(t, updatedBook.LoanedTo)
				assert.True(t, updatedBook.LoanedUntil.IsZero())
			})

		err := service.ReturnBook(context.Background(), "123")
		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})
}
