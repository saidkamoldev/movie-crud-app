package config

import (
	"go.uber.org/fx"
	"movie-crud-app/internal/delivery/handler"
	"movie-crud-app/internal/repository"
)

// NewApp - FX
func NewApp() *fx.App {
	return fx.New(
		fx.Provide(
			repository.ConnectDB,
			handler.NewMovieHandler, // handler.NewMovieHandler  provide
			NewRouter,               // NewRouter  provide
		),
	)
}
