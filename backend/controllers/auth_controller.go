package controllers

import (
	"log"
	"movie-app/models"
	"movie-app/repositories"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var JwtKey = []byte("SqYlIYQNBrUFu29bW8K4cCWarKiQetVKw2cUwx8pMSg=")

type Credentials struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

type LoginCredentials struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type AuthController struct {
	UserRepo *repositories.UserRepository
}

func NewAuthController(repo *repositories.UserRepository) *AuthController {
	return &AuthController{UserRepo: repo}
}

func (ac *AuthController) Signup(c *gin.Context) {
	var creds Credentials
	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(creds.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error hashing password:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	user := &models.User{
		Username:     creds.Username,
		PasswordHash: string(hashedPassword),
		Email:        creds.Email,
	}

	if err := ac.UserRepo.CreateUser(user); err != nil {
		log.Println("Error creating user:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created"})
}

func (ac *AuthController) Login(c *gin.Context) {
	var creds LoginCredentials
	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	user, err := ac.UserRepo.GetUserByUsername(creds.Username)
	if err != nil || bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(creds.Password)) != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Username: creds.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JwtKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
