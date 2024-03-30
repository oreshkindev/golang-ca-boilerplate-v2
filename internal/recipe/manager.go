package recipe

import (
	"github.com/oreshkindev/golang-ca-boilerplate-v2/internal/recipe/controller"
	"github.com/oreshkindev/golang-ca-boilerplate-v2/internal/recipe/repository"
	"github.com/oreshkindev/golang-ca-boilerplate-v2/internal/recipe/usecase"
	"github.com/oreshkindev/golang-ca-boilerplate-v2/pkg/postgres"
)

type Manager struct {
	Repository repository.Repository
	Usecase    usecase.Usecase
	Controller controller.Controller
}

func New(postgres *postgres.Database) *Manager {
	repository := repository.New(postgres)
	usecase := usecase.New(repository)
	controller := controller.New(usecase)

	return &Manager{
		Repository: *repository,
		Usecase:    *usecase,
		Controller: *controller,
	}
}
