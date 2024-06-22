package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Movie struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Poster string `json:"poster_path"`
}

func FetchMovieFromTMDB(movieID string) (*Movie, error) {
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

	var movie Movie
	if err := json.NewDecoder(resp.Body).Decode(&movie); err != nil {
		return nil, err
	}

	return &movie, nil
}
