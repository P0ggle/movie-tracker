package controllers

import (
	"movie-app/models"
	"movie-app/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MovieController struct {
	MovieService *services.MovieService
}

func NewMovieController(movieService *services.MovieService) *MovieController {
	return &MovieController{MovieService: movieService}
}

func (mc *MovieController) GetMovie(c *gin.Context) {
	movieID := c.Param("id")
	movie, err := mc.MovieService.FetchMovieFromTMDB(movieID)
	if err != nil {
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

	movies, err := mc.MovieService.SearchByName(name)
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

	username := c.MustGet("username").(string)
	request.Username = username

	err := mc.MovieService.AddMovieToList(&request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, request)
}

func (mc *MovieController) GetMoviesToWatch(c *gin.Context) {
	username := c.MustGet("username").(string)

	movies, err := mc.MovieService.GetMoviesToWatch(username)
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	movieID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid movie ID"})
		return
	}

	err = mc.MovieService.UpdateWatchedStatus(movieID, *request.Watched)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Watched status updated"})
}
