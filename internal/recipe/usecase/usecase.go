package usecase

import "github.com/oreshkindev/golang-ca-boilerplate-v2/internal/recipe/entity"

type Usecase struct {
	repository entity.Repository
}

func New(repository entity.Repository) *Usecase {
	return &Usecase{repository: repository}
}

func (usecase *Usecase) Get() (string, error) {
	return usecase.repository.Get()
}
