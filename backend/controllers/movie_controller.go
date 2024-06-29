package controllers

import (
	"database/sql"
	"log"
	"movie-app/models"
	"movie-app/services"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type AddMovieRequest struct {
	Name       string `json:"name" binding:"required"`
	PosterPath string `json:"poster_path" binding:"required"`
}

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

func AddMovieToList(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request AddMovieRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		movie := models.MovieToWatch{
			Name:       request.Name,
			PosterPath: request.PosterPath,
			TimeAdded:  time.Now(),
		}

		query := `INSERT INTO movies_to_watch (name, poster_path, time_added) VALUES ($1, $2, $3) RETURNING id`
		err := db.QueryRow(query, movie.Name, movie.PosterPath, movie.TimeAdded).Scan(&movie.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, movie)
	}
}

func GetMoviesToWatch(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		rows, err := db.Query("SELECT id, name, poster_path, time_added FROM movies_to_watch")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer rows.Close()

		var movies []models.MovieToWatch
		for rows.Next() {
			var movie models.MovieToWatch
			if err := rows.Scan(&movie.ID, &movie.Name, &movie.PosterPath, &movie.TimeAdded); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			movies = append(movies, movie)
		}

		if err := rows.Err(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, movies)
	}
}
