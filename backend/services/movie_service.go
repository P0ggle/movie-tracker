package services

import (
	"encoding/json"
	"fmt"
	"io"
	"movie-app/models"
	"movie-app/repositories"
	"net/http"
	"os"
	"strings"
	"time"
)

type MovieService struct {
	MovieRepo *repositories.MovieRepository
	UserRepo  *repositories.UserRepository
}

func NewMovieService(movieRepo *repositories.MovieRepository, userRepo *repositories.UserRepository) *MovieService {
	return &MovieService{MovieRepo: movieRepo, UserRepo: userRepo}
}

func (ms *MovieService) FetchMovieFromTMDB(movieID string) (*models.Movie, error) {
	apiKey := os.Getenv("TMDB_API_KEY")
	url := fmt.Sprintf("https://api.themoviedb.org/3/movie/%s?api_key=%s", movieID, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: unable to fetch movie data, status code: %d", resp.StatusCode)
	}

	var movie models.Movie
	if err := json.NewDecoder(resp.Body).Decode(&movie); err != nil {
		return nil, err
	}

	return &movie, nil
}

func (ms *MovieService) SearchByName(name string) ([]models.Movie, error) {
	apiKey := os.Getenv("TMDB_API_KEY")
	name = strings.ReplaceAll(strings.TrimSpace(name), " ", ",")

	url := fmt.Sprintf("https://api.themoviedb.org/3/search/multi?api_key=%s&query=%s&include_adult=false&language=en-US&page=1", apiKey, name)
	resp, err := http.DefaultClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: unable to fetch movie data, status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var searchResults models.SearchResults
	if err := json.Unmarshal(body, &searchResults); err != nil {
		return nil, err
	}

	if len(searchResults.Results) == 0 {
		return nil, fmt.Errorf("no results found")
	}

	var filteredResults []models.Movie
	for _, result := range searchResults.Results {
		if result.OriginalLanguage == "en" {
			filteredResults = append(filteredResults, result)
		}
	}

	if len(filteredResults) > 10 {
		return filteredResults[:10], nil
	}

	return filteredResults, nil
}

func (ms *MovieService) AddMovieToList(request *models.MovieToWatch) error {
	user, err := ms.UserRepo.GetUserByUsername(request.Username)
	if err != nil {
		return err
	}

	request.TimeAdded = time.Now()
	request.UserID = user.ID
	return ms.MovieRepo.AddMovie(request)
}

func (ms *MovieService) GetMoviesToWatch(username string) ([]models.MovieToWatch, error) {
	user, err := ms.UserRepo.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}

	return ms.MovieRepo.GetMoviesByUser(user.ID)
}

func (ms *MovieService) UpdateWatchedStatus(movieID int, watched bool) error {
	rowsAffected, err := ms.MovieRepo.UpdateWatchedStatus(movieID, watched)
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("movie not found")
	}

	return nil
}
