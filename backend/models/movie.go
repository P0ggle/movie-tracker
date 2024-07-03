package models

import "time"

type Movie struct {
	ID            int    `json:"id"`
	OriginalTitle string `json:"original_title,omitempty"`
	OriginalName  string `json:"original_name,omitempty"`
	PosterPath    string `json:"poster_path"`
	Overview      string `json:"overview"`
}

type MovieToWatch struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	PosterPath string    `json:"poster_path"`
	TimeAdded  time.Time `json:"time_added"`
	Watched    bool      `json:"watched"`
}

type SearchResults struct {
	Results []Movie `json:"results"`
}
