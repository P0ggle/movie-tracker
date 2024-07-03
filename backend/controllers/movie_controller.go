package controllers

import (
	"log"
	"movie-app/models"
	"movie-app/repositories"
	"movie-app/services"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type MovieController struct {
	MovieRepo *repositories.MovieRepository
}

func NewMovieController(repo *repositories.MovieRepository) *MovieController {
	return &MovieController{MovieRepo: repo}
}

func (mc *MovieController) GetMovie(c *gin.Context) {
	movieID := c.Param("id")
	movie, err := services.FetchMovieFromTMDB(movieID)
	if err != nil {
		log.Println("Error fetching movie:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, movie)
}

func (mc *MovieController) SearchContent(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Name parameter is required"})
		return
	}

	movies, err := services.SearchByName(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, movies)
}

func (mc *MovieController) AddMovieToList(c *gin.Context) {
	var request models.MovieToWatch
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	request.TimeAdded = time.Now()
	if err := mc.MovieRepo.AddMovie(&request); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, request)
}

func (mc *MovieController) GetMoviesToWatch(c *gin.Context) {
	movies, err := mc.MovieRepo.GetMovies()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, movies)
}

func (mc *MovieController) UpdateWatchedStatus(c *gin.Context) {
	var request struct {
		Watched *bool `json:"watched" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		log.Printf("Bind JSON error: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Printf("Parsed request: %+v\n", request)

	movieID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println("Invalid movie ID:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid movie ID"})
		return
	}

	rowsAffected, err := mc.MovieRepo.UpdateWatchedStatus(movieID, *request.Watched)
	if err != nil {
		log.Println("DB error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if rowsAffected == 0 {
		log.Println("No rows affected")
		c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Watched status updated"})
}

