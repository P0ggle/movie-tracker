package routes

import (
	"database/sql"
	"movie-app/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, db *sql.DB) {
	router.GET("/search", controllers.SearchContent)
	router.GET("/movies/:id", controllers.GetMovie)
	router.POST("/add-to-list", controllers.AddMovieToList(db))
	router.GET("/movies-to-watch", controllers.GetMoviesToWatch(db))
}
