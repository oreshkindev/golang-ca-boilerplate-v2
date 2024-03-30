package internal

import (
	"github.com/oreshkindev/golang-ca-boilerplate-v2/internal/post"
	"github.com/oreshkindev/golang-ca-boilerplate-v2/internal/recipe"
	"github.com/oreshkindev/golang-ca-boilerplate-v2/pkg/postgres"
)

type Manager struct {
	Post   *post.Manager
	Recipe *recipe.Manager
}

func New(postgres *postgres.Database) *Manager {
	return &Manager{
		Post:   post.New(postgres),
		Recipe: recipe.New(postgres),
	}
}
