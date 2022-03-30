package commands

import (
	"context"
	"fmt"

	"github.com/neo-classic/golang-cqrs-mvp/internal/book/domain"
)

type BookRepo interface {
	Add(ctx context.Context, book domain.Book) error
}

type AddBook struct {
	Book domain.Book
}

type AddBookHandler struct {
	repo BookRepo
}

func NewAddBookHandler(repo BookRepo) AddBookHandler {
	return AddBookHandler{
		repo: repo,
	}
}

func (h AddBookHandler) Handle(ctx context.Context, cmd AddBook) (domain.BookID, error) {
	if !cmd.Book.ISBN.IsValid() {
		return "", fmt.Errorf("invalid ISBN number")
	}

	err := h.repo.Add(ctx, cmd.Book)
	if err != nil {
		return "", err
	}

	return cmd.Book.ID, nil
}
