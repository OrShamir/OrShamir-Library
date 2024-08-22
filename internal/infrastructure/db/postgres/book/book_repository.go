package postgres

import (
	"Or/Library/internal/domain/book"
	"context"
	"errors"
	"net/url"
	"regexp"

	"gorm.io/gorm"
)

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) book.BookRepository {
	return &bookRepository{db}
}

func (r *bookRepository) Create(ctx context.Context, book *book.Book) error {
	entity := MapToBookEntity(book)
	return r.db.WithContext(ctx).Create(&entity).Error
}

func (r *bookRepository) GetByID(ctx context.Context, id string) (*book.Book, error) {
	var entity BookEntity
	err := r.db.WithContext(ctx).First(&entity, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("book not found")
		}
		return nil, err
	}
	return MapToBookDomain(&entity), nil
}

func (r *bookRepository) Update(ctx context.Context, book *book.Book) error {
	entity := MapToBookEntity(book)
	return r.db.WithContext(ctx).Save(&entity).Error
}

func (r *bookRepository) Delete(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Delete(&BookEntity{}, "id = ?", id).Error
}

func (r *bookRepository) Search(ctx context.Context, query string) ([]*book.Book, error) {
	if !isValidInput(query) {
		return nil, errors.New("invalid input")
	}

	sanitizedQuery := sanitizeInput(query)

	var entities []BookEntity
	err := r.db.WithContext(ctx).
		Where(BookEntity{Title: sanitizedQuery}).
		Or(BookEntity{Author: sanitizedQuery}).
		Or(BookEntity{Topic: sanitizedQuery}).
		Find(&entities).Error
	if err != nil {
		return nil, err
	}

	books := make([]*book.Book, len(entities))
	for i, entity := range entities {
		books[i] = MapToBookDomain(&entity)
	}
	return books, nil
}

func sanitizeInput(input string) string {
	return url.QueryEscape(input)
}

func isValidInput(input string) bool {
	isAlphanumeric := regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString
	return isAlphanumeric(input)
}
