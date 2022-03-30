package mongo

import (
	"context"
	"fmt"

	"github.com/neo-classic/golang-cqrs-mvp/internal/book/domain"
)

type BookRepo struct {
	books []domain.Book
}

func NewBookRepo() *BookRepo {
	return &BookRepo{}
}

func (r *BookRepo) Add(ctx context.Context, book domain.Book) error {
	r.books = append(r.books, book)
	return nil
}

func (r *BookRepo) LoadByID(ctx context.Context, id domain.BookID) (*domain.Book, error) {
	for _, b := range r.books {
		if b.ID == id {
			return &b, nil
		}
	}
	return nil, fmt.Errorf("book not found")
}
