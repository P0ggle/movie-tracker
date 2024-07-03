package routes

import (
	"database/sql"
	"log"
	"movie-app/controllers"
	"movie-app/repositories"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, db *sql.DB) {
	movieRepo := &repositories.MovieRepository{DB: db}
	userRepo := &repositories.UserRepository{DB: db}
	movieController := controllers.NewMovieController(movieRepo, userRepo)
	authController := controllers.NewAuthController(userRepo)

	router.POST("/signup", authController.Signup)
	router.POST("/login", authController.Login)
	router.GET("/search", movieController.SearchContent)

	auth := router.Group("/")
	auth.Use(AuthMiddleware())

	auth.GET("/movies/:id", movieController.GetMovie)
	auth.POST("/add-to-list", movieController.AddMovieToList)
	auth.GET("/movies-to-watch", movieController.GetMoviesToWatch)
	auth.PUT("/movies/:id/watched", movieController.UpdateWatchedStatus)
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		if len(tokenString) > 7 && strings.ToUpper(tokenString[:7]) == "BEARER " {
			tokenString = tokenString[7:]
		}

		claims := &controllers.Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return controllers.JwtKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		log.Printf("Authenticated user: %s", claims.Username)

		c.Set("username", claims.Username)
		c.Next()
	}
}
