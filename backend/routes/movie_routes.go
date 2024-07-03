package routes

import (
	"database/sql"
	"movie-app/controllers"
	"movie-app/repositories"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, db *sql.DB) {
	movieRepo := &repositories.MovieRepository{DB: db}
	movieController := controllers.NewMovieController(movieRepo)

	router.GET("/search", movieController.SearchContent)
	router.GET("/movies/:id", movieController.GetMovie)
	router.POST("/add-to-list", movieController.AddMovieToList)
	router.GET("/movies-to-watch", movieController.GetMoviesToWatch)
	router.PUT("/movies/:id/watched", movieController.UpdateWatchedStatus)
}
