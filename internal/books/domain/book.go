package domain

import (
	"errors"

	"github.com/google/uuid"
)

var ErrorISBNIsNotValid = errors.New("ISBN is not valid")

type Book struct {
	ID   BookID
	ISBN ISBNNuber
	Name BookName
}

func NewBook(isbn ISBNNuber, name BookName) (*Book, error) {
	if isbn.IsValid() {
		return &Book{
			ID:   NewBookID(),
			ISBN: isbn,
			Name: name,
		}, nil
	}

	return nil, ErrorISBNIsNotValid
}

type BookID string

func NewBookID() BookID {
	return BookID(uuid.New().String())
}

type ISBNNuber struct {
	Prefix        int
	CountryCode   int
	PublisherCode int
	BookCode      int
	ControlNumber int
}

func (isbn *ISBNNuber) IsValid() bool {
	if isbn.Prefix < 0 || isbn.Prefix > 999 {
		return false
	}

	if isbn.CountryCode < 0 || isbn.CountryCode > 200 {
		return false
	}

	if isbn.PublisherCode < 0 || isbn.PublisherCode > 99999 {
		return false
	}

	if isbn.BookCode < 0 || isbn.BookCode > 999 {
		return false
	}

	if isbn.ControlNumber < 0 || isbn.ControlNumber > 9 {
		return false
	}

	return true
}

type BookName string

func (s BookName) String() string {
	return string(s)
}
