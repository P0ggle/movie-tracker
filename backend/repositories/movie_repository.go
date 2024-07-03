package repositories

import (
	"database/sql"
	"log"
	"movie-app/models"
)

type MovieRepository struct {
	DB *sql.DB
}

func (r *MovieRepository) AddMovie(movie *models.MovieToWatch) error {
	query := `INSERT INTO movies_to_watch (name, poster_path, time_added, user_id) VALUES ($1, $2, $3, $4) RETURNING id`
	err := r.DB.QueryRow(query, movie.Name, movie.PosterPath, movie.TimeAdded, movie.UserID).Scan(&movie.ID)
	if err != nil {
		log.Printf("Error executing query: %v\n", err)
		return err
	}
	return nil
}

func (r *MovieRepository) GetMoviesByUser(userID int) ([]models.MovieToWatch, error) {
	rows, err := r.DB.Query("SELECT id, name, poster_path, time_added, watched FROM movies_to_watch WHERE user_id = $1", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []models.MovieToWatch
	for rows.Next() {
		var movie models.MovieToWatch
		if err := rows.Scan(&movie.ID, &movie.Name, &movie.PosterPath, &movie.TimeAdded, &movie.Watched); err != nil {
			return nil, err
		}
		movies = append(movies, movie)
	}

	return movies, rows.Err()
}

func (r *MovieRepository) UpdateWatchedStatus(movieID int, watched bool) (int64, error) {
	query := `UPDATE movies_to_watch SET watched = $1 WHERE id = $2`
	result, err := r.DB.Exec(query, watched, movieID)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}
