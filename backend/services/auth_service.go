package services

import (
	"movie-app/models"
	"movie-app/repositories"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

var JwtKey = []byte("SqYlIYQNBrUFu29bW8K4cCWarKiQetVKw2cUwx8pMSg=")

type AuthService struct {
	UserRepo *repositories.UserRepository
}

func NewAuthService(userRepo *repositories.UserRepository) *AuthService {
	return &AuthService{UserRepo: userRepo}
}

func (as *AuthService) Signup(creds *models.Credentials) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(creds.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &models.User{
		Username:     creds.Username,
		PasswordHash: string(hashedPassword),
		Email:        creds.Email,
	}

	return as.UserRepo.CreateUser(user)
}

func (as *AuthService) Login(creds *models.LoginCredentials) (string, error) {
	user, err := as.UserRepo.GetUserByUsername(creds.Username)
	if err != nil || bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(creds.Password)) != nil {
		return "", err
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &models.Claims{
		Username: creds.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
