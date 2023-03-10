package repository

import (
	"github.com/jmoiron/sqlx"
	"music-player/pkg/model"
	"time"
)

type Player interface {
	// CRUD
	AddSong(title string, duration time.Duration) (int, error)
	UpdateSong(playlistId, songId int, input model.UpdateSongInput) error
	GetSongs(playlistId int) ([]model.Song, error)
	DeleteSong(songId int) error
}

type Repository struct {
	Player
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Player: NewPlayerPostgres(db),
	}
}
