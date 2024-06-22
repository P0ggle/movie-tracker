package routes

import (
	"movie-app/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterMovieRoutes(router *gin.Engine) {
	router.GET("/search", controllers.SearchMovie)
	router.GET("/movies/:id", controllers.GetMovie)
}
