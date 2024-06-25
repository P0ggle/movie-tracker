package controllers

import (
	"log"
	"movie-app/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetMovie(c *gin.Context) {
	movieID := c.Param("id")
	movie, err := services.FetchMovieFromTMDB(movieID)
	if err != nil {
		log.Println("Error fetching movie:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, movie)
}

func SearchContent(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "name query parameter is required"})
		return
	}

	content, err := services.SearchByName(name)
	if err != nil {
		log.Println("Error searching content:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, content)
}
