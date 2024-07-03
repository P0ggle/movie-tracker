package repositories

import (
	"database/sql"
	"movie-app/models"
)

type UserRepository struct {
	DB *sql.DB
}

func (r *UserRepository) CreateUser(user *models.User) error {
	query := `INSERT INTO users (username, password_hash, email) VALUES ($1, $2, $3) RETURNING id, created_at`
	return r.DB.QueryRow(query, user.Username, user.PasswordHash, user.Email).Scan(&user.ID, &user.CreatedAt)
}

func (r *UserRepository) GetUserByUsername(username string) (*models.User, error) {
	query := `SELECT id, username, password_hash, email, created_at FROM users WHERE username = $1`
	user := &models.User{}
	err := r.DB.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.PasswordHash, &user.Email, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return user, nil
}
