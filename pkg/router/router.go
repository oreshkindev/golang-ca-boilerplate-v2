package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/oreshkindev/golang-ca-boilerplate-v2/internal"
)

type Router struct {
	*chi.Mux
}

func NewRouter(manager *internal.Manager) (*Router, error) {
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	}))

	router.Use(render.SetContentType(render.ContentTypeJSON))

	router.Route("/v1", func(r chi.Router) {
		r.Mount("/posts", postsHandler(manager))
		r.Mount("/recipes", recipesHandler(manager))
	})

	return &Router{router}, nil
}

func postsHandler(manager *internal.Manager) http.Handler {
	router := chi.NewRouter()
	router.Get("/", manager.Post.Controller.Get)

	return router
}

func recipesHandler(manager *internal.Manager) http.Handler {
	router := chi.NewRouter()
	router.Get("/", manager.Recipe.Controller.Get)

	return router
}
