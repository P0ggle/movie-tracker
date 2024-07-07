package services

import (
	"encoding/json"
	"fmt"
	"io"
	"movie-app/models"
	"net/http"
	"os"
	"strings"
)

func FetchMovieFromTMDB(movieID string) (*models.Movie, error) {
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

func SearchByName(name string) ([]models.Movie, error) {
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

	// Filter results to include only English-language movies and TV shows
	var filteredResults []models.Movie
	for _, result := range searchResults.Results {
		if result.OriginalLanguage == "en" {
			filteredResults = append(filteredResults, result)
		}
	}

	// Return only the first 10 filtered results if there are more than 10
	if len(filteredResults) > 10 {
		return filteredResults[:10], nil
	}

	return filteredResults, nil
}
