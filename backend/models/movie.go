package models

import "time"

type Movie struct {
	ID               int    `json:"id"`
	OriginalTitle    string `json:"original_title,omitempty"`
	OriginalName     string `json:"original_name,omitempty"`
	PosterPath       string `json:"poster_path"`
	Overview         string `json:"overview"`
	OriginalLanguage string `json:"original_language"`
}

type MovieToWatch struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	PosterPath string    `json:"poster_path"`
	TimeAdded  time.Time `json:"time_added"`
	Watched    bool      `json:"watched"`
	UserID     int       `json:"user_id"`
	Username   string    `json:"username"`
}

type SearchResults struct {
	Results []Movie `json:"results"`
}
