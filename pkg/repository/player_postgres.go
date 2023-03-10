package repository

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"music-player/pkg/model"
	"strings"
	"time"
)

type PlayerPostgres struct {
	db *sqlx.DB
}

func (r *PlayerPostgres) UpdateSong(playlistId, songId int, input model.UpdateSongInput) error {

	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.Duration != nil {
		setValues = append(setValues, fmt.Sprintf("duration=$%d", argId))
		args = append(args, *input.Duration)
		argId++
	}

	// title=$1
	// duration=$1
	// title=$1, duration=$2
	setQuery := strings.Join(setValues, ", ")

	tx, err := r.db.Begin()

	if err != nil {
		return err
	}

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=$%d", songsTable, setQuery, argId, argId+1)
	args = append(args, playlistId, songId)

	logrus.Warnf("updateQuery: %s", query)
	logrus.Warnf("args: %s", args)

	_, err = tx.Exec(query, args...)

	if err != nil {
		logrus.Warnf("failed to update song (%d), cause: %s", err.Error())
		tx.Rollback()
		return err
	}

	return err

}

func (r *PlayerPostgres) GetSongs(playlistId int) ([]model.Song, error) {
	var songs []model.Song

	// playlistId нужен будет при создании нескольких плейлистов
	query := fmt.Sprintf("SELECT songs.title, songs.duration FROM %s songs", songsTable)
	err := r.db.Select(&songs, query)
	return songs, err
}

func (r *PlayerPostgres) DeleteSong(songId int) error {
	query := fmt.Sprintf("DELETE FROM %s s WHERE s.id=$1", songsTable)
	result, err := r.db.Exec(query, songId)

	deletedRows, _ := result.RowsAffected()

	if deletedRows == 0 {
		return errors.New(fmt.Sprintf("cannot delete songId=%d", songId))
	}

	logrus.Infof("successfully deleted")

	return err
}

func (r *PlayerPostgres) AddSong(title string, duration time.Duration) (int, error) {

	playlistId := 1 // default

	tx, err := r.db.Begin()

	if err != nil {
		return 0, err
	}

	var songId int

	createPostQuery := fmt.Sprintf("INSERT INTO %s (playlist_id, title, description) VALUES ($1, $2, $3) RETURNING id", songsTable)
	row := tx.QueryRow(createPostQuery, songId, playlistId, title, duration)

	if err := row.Scan(&songId); err != nil {
		logrus.Errorf("[%s] [DB] failed to insert song (%d) into playlist (%d)",
			time.Now().UTC().Format("2006-01-02 15:04:05"), songId, playlistId)
		tx.Rollback()
		return 0, err
	}

	logrus.Infof("successfully inserted into db")

	return songId, tx.Commit()
}

func NewPlayerPostgres(db *sqlx.DB) *PlayerPostgres {
	return &PlayerPostgres{db: db}
}
