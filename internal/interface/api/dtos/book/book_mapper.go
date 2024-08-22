package book

import "Or/Library/internal/domain/book"

func MapToBookDomainFromDTO(dto *BookDTO) *book.Book {
	return &book.Book{
		ID:         dto.ID,
		Title:      dto.Title,
		Author:     dto.Author,
		Topic:      dto.Topic,
		Year:       dto.Year,
		Popularity: dto.Popularity,
	}
}

func MapToBookDTO(b *book.Book) BookDTO {
	return BookDTO{
		ID:         b.ID,
		Title:      b.Title,
		Author:     b.Author,
		Topic:      b.Topic,
		Year:       b.Year,
		Popularity: b.Popularity,
	}
}
