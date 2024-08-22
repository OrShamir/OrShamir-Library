package book

import (
	"context"
)

type BookRepository interface {
	Create(ctx context.Context, book *Book) error
	GetByID(ctx context.Context, id string) (*Book, error)
	Update(ctx context.Context, book *Book) error
	Delete(ctx context.Context, id string) error
	Search(ctx context.Context, query string) ([]*Book, error)
}
