package main

import (
	"net/http"
	"time"

	"github.com/oreshkindev/golang-ca-boilerplate-v2/internal"
	"github.com/oreshkindev/golang-ca-boilerplate-v2/pkg/postgres"
	"github.com/oreshkindev/golang-ca-boilerplate-v2/pkg/router"
)

const (
	dbURL        = "postgres://postgres@localhost/postgres?sslmode=disable"
	maxPoolSize  = 5
	connAttempts = 10
	connTimeout  = 5 * time.Second
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	postgres, err := postgres.New(dbURL, postgres.MaxPoolSize(maxPoolSize), postgres.ConnAttempts(connAttempts), postgres.ConnTimeout(connTimeout))
	if err != nil {
		return err
	}
	defer postgres.Close()

	manager := internal.New(postgres)

	router, err := router.NewRouter(manager)
	if err != nil {
		return err
	}

	return http.ListenAndServe(":9000", router)
}
