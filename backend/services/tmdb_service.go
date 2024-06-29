package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type Movie struct {
	ID            int    `json:"id"`
	OriginalTitle string `json:"original_title,omitempty"`
	OriginalName  string `json:"original_name,omitempty"`
	PosterPath    string `json:"poster_path"`
	Overview      string `json:"overview"`
}

type SearchResults struct {
	Results []Movie `json:"results"`
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

func SearchByName(name string) ([]Movie, error) {
	apiKey := os.Getenv("TMDB_API_KEY")
	name = strings.ReplaceAll(strings.TrimSpace(name), " ", ",")

	url := fmt.Sprintf("https://api.themoviedb.org/3/search/multi?api_key=%s&query=%s", apiKey, name)
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

	var searchResults SearchResults
	if err := json.Unmarshal(body, &searchResults); err != nil {
		return nil, err
	}

	if len(searchResults.Results) == 0 {
		return nil, fmt.Errorf("no results found")
	}

	return searchResults.Results, nil
}
