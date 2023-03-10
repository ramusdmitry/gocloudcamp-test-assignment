package model

import "time"

type Song struct {
	Title    string
	Duration time.Duration
}

type UpdateSongInput struct {
	Title    *string        `json:"title"`
	Duration *time.Duration `json:"duration"`
}
