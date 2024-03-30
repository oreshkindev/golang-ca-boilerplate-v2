package repository

import "github.com/oreshkindev/golang-ca-boilerplate-v2/pkg/postgres"

type Repository struct {
	database *postgres.Database
}

func New(database *postgres.Database) *Repository {
	return &Repository{database: database}
}

func (repository *Repository) Get() (string, error) {

	return "result from Post repository", nil
}
