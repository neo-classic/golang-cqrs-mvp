package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/neo-classic/golang-cqrs-mvp/internal/book/application"
	"github.com/neo-classic/golang-cqrs-mvp/internal/book/application/commands"
	"github.com/neo-classic/golang-cqrs-mvp/internal/book/domain"
)

type BookHTTPHandler struct {
	r       *gin.Engine
	service application.ApplicationService
}

func NewBookHTTPHandler(r *gin.Engine, service application.ApplicationService) *BookHTTPHandler {
	h := &BookHTTPHandler{
		r:       r,
		service: service,
	}
	r.Handle(http.MethodPost, "/book", h.Add)
	r.Handle(http.MethodGet, "/book/:id", h.FindByID)

	return h
}

func (h *BookHTTPHandler) Add(c *gin.Context) {
	dto := new(BookRequest)
	err := c.BindJSON(dto)
	if err != nil {
		return
	}

	book, err := domain.NewBook(dto.toISBNDomain(), domain.BookName(dto.Name))
	if err != nil {
		return
	}

	bookID, err := h.service.Commands.AddBook.Handle(c.Request.Context(), commands.AddBook{
		Book: *book,
	})
	if err != nil {
		return
	}

	c.JSON(http.StatusCreated, AddBookResponse{
		ID: bookID.String(),
	})
}

func (h *BookHTTPHandler) FindByID(c *gin.Context) {

}
