package controllers

import (
	"movie-app/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetMovie(c *gin.Context) {
	movieID := c.Param("id")
	movie, err := services.FetchMovieFromTMDB(movieID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, movie)
}
