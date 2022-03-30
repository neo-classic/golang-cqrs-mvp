package application

import (
	"github.com/neo-classic/golang-cqrs-mvp/internal/book/application/commands"
	"github.com/neo-classic/golang-cqrs-mvp/internal/book/application/queries"
)

type ApplicationService struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	AddBook commands.AddBookHandler
}

type Queries struct {
	GetBook queries.GetBookHandler
}

func NewApplicationService(commands Commands, queries Queries) ApplicationService {
	return ApplicationService{
		Commands: commands,
		Queries:  queries,
	}
}
