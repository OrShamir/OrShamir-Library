package postgres_test

import (
	"Or/Library/internal/domain/book"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	bookDb "Or/Library/internal/infrastructure/db/postgres/book"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
	"testing"
	"time"
)

func setupDatabase() (*gorm.DB, func()) {
	// Use sqlite for testing purposes
	database, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	// AutoMigrate our Product model
	err = database.AutoMigrate(&bookDb.BookEntity{})
	if err != nil {
		panic("Failed to migrate database")
	}

	// Cleanup function to truncate tables
	cleanup := func() {
		database.Exec("DELETE FROM booka")
	}

	return database, cleanup

}

func TestBookRepository_Create_ShouldSuccess(t *testing.T) {
	db, cleanup := setupDatabase()
	defer cleanup()

	bookRepository := bookDb.NewBookRepository(db)
	err := bookRepository.Create(context.Background(), &book.Book{
		"1", "The Hobbit", "J.R.R. Tolkien", "Fantasy",
		1937, 5, false, "", time.Time{}})
	assert.NoError(t, err)
}

// TestBookRepository_GetByID_ShouldSuccess tests the GetByID method of the BookRepository

// ... more tests cases
