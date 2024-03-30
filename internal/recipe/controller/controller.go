package controller

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/oreshkindev/golang-ca-boilerplate-v2/internal/recipe/entity"
)

type Controller struct {
	usecase entity.Usecase
}

func New(usecase entity.Usecase) *Controller {
	return &Controller{usecase: usecase}
}

func (controller *Controller) Get(w http.ResponseWriter, r *http.Request) {

	result, err := controller.usecase.Get()
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	render.JSON(w, r, result)
}
