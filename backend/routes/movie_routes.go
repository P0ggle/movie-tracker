package routes

import (
	"movie-app/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterMovieRoutes(router *gin.Engine) {
	router.GET("/search", controllers.SearchContent)
	router.GET("/movies/:id", controllers.GetMovie)
}
