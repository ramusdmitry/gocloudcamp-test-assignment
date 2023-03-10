package service

import (
	"music-player/pkg/model"
	"music-player/pkg/repository"
)

type PlayService struct {
	repo repository.Player
}

func (s *PlayService) AddSong(song *model.Song) (int, error) {
	return s.repo.AddSong(song.Title, song.Duration)
}

func (s *PlayService) UpdateSong(playlistId, songId int, input model.UpdateSongInput) error {
	return s.repo.UpdateSong(playlistId, songId, input)
}

func (s *PlayService) GetSongs(playlistId int) ([]model.Song, error) {
	return s.repo.GetSongs(playlistId)
}

func (s *PlayService) DeleteSong(songId int) error {

	// дописать проверку на случай, если в данный момент трек играет

	return s.repo.DeleteSong(songId)
}

func (p PlayService) Play() {
	//TODO implement me
	panic("implement me")
}

func (p PlayService) Pause() {
	//TODO implement me
	panic("implement me")
}

func (p PlayService) Next() error {
	//TODO implement me
	panic("implement me")
}

func (p PlayService) Prev() error {
	//TODO implement me
	panic("implement me")
}

func NewPlayService(repo repository.Player) *PlayService {
	return &PlayService{
		repo: repo,
	}
}
