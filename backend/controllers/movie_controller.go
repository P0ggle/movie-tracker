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
	UserRepo  *repositories.UserRepository
}

func NewMovieController(movieRepo *repositories.MovieRepository, userRepo *repositories.UserRepository) *MovieController {
	return &MovieController{MovieRepo: movieRepo, UserRepo: userRepo}
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
		log.Printf("Error binding JSON: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	username := c.MustGet("username").(string)
	user, err := mc.UserRepo.GetUserByUsername(username) // Use UserRepository here
	if err != nil {
		log.Printf("Error fetching user: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch user"})
		return
	}

	request.TimeAdded = time.Now()
	request.UserID = user.ID // Add User ID to the request
	if err := mc.MovieRepo.AddMovie(&request); err != nil {
		log.Printf("Error adding movie to repository: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("Movie added successfully: %+v\n", request)
	c.JSON(http.StatusOK, request)
}

func (mc *MovieController) GetMoviesToWatch(c *gin.Context) {
	username := c.MustGet("username").(string)
	user, err := mc.UserRepo.GetUserByUsername(username) // Use UserRepository here
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch user"})
		return
	}

	movies, err := mc.MovieRepo.GetMoviesByUser(user.ID)
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
