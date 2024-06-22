package main

import (
	"database/sql"
	"log"
	"movie-app/routes"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	dbURL := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := gin.Default()

	// Apply the CORS middleware
	r.Use(cors.Default())

	routes.RegisterMovieRoutes(r)
	r.Run()
}
