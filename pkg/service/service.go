package service

import (
	"music-player/pkg/model"
	"music-player/pkg/repository"
)

type Player interface {
	AddSong(song *model.Song) (int, error)
	UpdateSong(playlistId, songId int, input model.UpdateSongInput) error
	GetSongs(playlistId int) ([]model.Song, error)
	DeleteSong(songId int) error
	Play()
	Pause()
	Next() error
	Prev() error
}

type Service struct {
	Player
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Player: NewPlayService(repos.Player),
	}
}
