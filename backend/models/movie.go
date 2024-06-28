package models

import "time"

type MovieToWatch struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	PosterPath string    `json:"poster_path"`
	TimeAdded  time.Time `json:"time_added"`
}
