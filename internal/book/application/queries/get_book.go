package queries

import (
	"context"

	"github.com/neo-classic/golang-cqrs-mvp/internal/book/domain"
)

type BookRepo interface {
	LoadByID(ctx context.Context, id domain.BookID) (*domain.Book, error)
}

type GetBook struct {
	ID domain.BookID
}

type GetBookHandler struct {
	repo BookRepo
}

func NewGetBookHandler(repo BookRepo) GetBookHandler {
	return GetBookHandler{
		repo: repo,
	}
}

func (h GetBookHandler) Handle(ctx context.Context, cmd GetBook) (*domain.Book, error) {
	return h.repo.LoadByID(ctx, cmd.ID)
}
