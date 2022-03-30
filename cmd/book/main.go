package main

import (
	"github.com/gin-gonic/gin"
	"github.com/neo-classic/golang-cqrs-mvp/internal/book/application"
	"github.com/neo-classic/golang-cqrs-mvp/internal/book/application/commands"
	"github.com/neo-classic/golang-cqrs-mvp/internal/book/application/queries"
	"github.com/neo-classic/golang-cqrs-mvp/internal/book/infrastructure/mongo"
	"github.com/neo-classic/golang-cqrs-mvp/internal/book/presentation/http"
)

func main() {
	router := gin.Default()

	repo := mongo.NewBookRepo()
	applicationService := application.NewApplicationService(
		application.Commands{
			AddBook: commands.NewAddBookHandler(repo),
		},
		application.Queries{
			GetBook: queries.NewGetBookHandler(repo),
		},
	)
	_ = http.NewBookHTTPHandler(router, applicationService)

	router.Run()
}
