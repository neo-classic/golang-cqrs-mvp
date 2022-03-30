package http

import (
	"github.com/neo-classic/golang-cqrs-mvp/internal/book/domain"
)

type BookRequest struct {
	ISBN ISBN   `json:"isbn"`
	Name string `json:"name"`
}

type ISBN struct {
	Prefix        int `json:"prefix"`
	CountryCode   int `json:"country_code"`
	PublisherCode int `json:"publisher_code"`
	BookCode      int `json:"book_code"`
	ControlNumber int `json:"control_number"`
}

func (r *BookRequest) toISBNDomain() domain.ISBNNumber {
	return domain.ISBNNumber{
		Prefix:        r.ISBN.Prefix,
		CountryCode:   r.ISBN.CountryCode,
		PublisherCode: r.ISBN.PublisherCode,
		BookCode:      r.ISBN.BookCode,
		ControlNumber: r.ISBN.ControlNumber,
	}
}

type AddBookResponse struct {
	ID string `json:"id"`
}
