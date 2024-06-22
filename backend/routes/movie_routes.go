package routes

import (
	"movie-app/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterMovieRoutes(r *gin.Engine) {
	r.GET("/movies/:id", controllers.GetMovie)
}
