package postgres

import (
	"Or/Library/internal/domain/book"
)

func MapToBookEntity(book *book.Book) *BookEntity {
	return &BookEntity{
		ID:          book.ID,
		Title:       book.Title,
		Author:      book.Author,
		Topic:       book.Topic,
		Year:        book.Year,
		Popularity:  book.Popularity,
		IsLoaned:    book.IsLoaned,
		LoanedTo:    book.LoanedTo,
		LoanedUntil: book.LoanedUntil,
	}
}

func MapToBookDomain(entity *BookEntity) *book.Book {
	return &book.Book{
		ID:          entity.ID,
		Title:       entity.Title,
		Author:      entity.Author,
		Topic:       entity.Topic,
		Year:        entity.Year,
		Popularity:  entity.Popularity,
		IsLoaned:    entity.IsLoaned,
		LoanedTo:    entity.LoanedTo,
		LoanedUntil: entity.LoanedUntil,
	}
}
