package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	swagger "github.com/swaggo/gin-swagger"
	"go.uber.org/fx"
	"movie-crud-app/config"

	_ "movie-crud-app/cmd/server/docs"

	"movie-crud-app/internal/delivery/handler"
	"movie-crud-app/internal/repository"
)

func registerRoutes(r *gin.Engine, movieHandler *handler.MovieHandler) {
	r.GET("/movies", movieHandler.GetMovies)
	r.POST("/movies", movieHandler.CreateMovie)
	r.PUT("/movies/:id", movieHandler.UpdateMovie)
	r.DELETE("/movies/:id", movieHandler.DeleteMovie)

	// Swagger UI-ni qo'shamiz
	r.GET("/swagger/*any", swagger.WrapHandler(swaggerFiles.Handler))
}

func main() {
	app := fx.New(
		fx.Provide(
			repository.ConnectDB,
			handler.NewMovieHandler,
			config.NewRouter,
		),
		fx.Invoke(registerRoutes),
	)

	app.Run()
}
