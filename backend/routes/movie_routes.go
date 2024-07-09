package routes

import (
	"database/sql"
	"movie-app/controllers"
	"movie-app/models" // Ensure this import is present
	"movie-app/repositories"
	"movie-app/services"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, db *sql.DB) {
	userRepo := &repositories.UserRepository{DB: db}
	movieRepo := &repositories.MovieRepository{DB: db}

	authService := services.NewAuthService(userRepo)
	movieService := services.NewMovieService(movieRepo, userRepo)

	authController := controllers.NewAuthController(authService)
	movieController := controllers.NewMovieController(movieService)

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

		claims := &models.Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return services.JwtKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Set("username", claims.Username)
		c.Next()
	}
}
