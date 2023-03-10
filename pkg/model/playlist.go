package model

import (
	"container/list"
	"errors"
	"fmt"
	"sync"
	"time"
)

type Playlist struct {
	songs        *list.List
	CurrentSong  *list.Element
	currentTimer *time.Timer
	paused       bool
	mutex        sync.Mutex
}

func NewPlaylist() *Playlist {
	return &Playlist{
		songs: list.New(),
	}
}

func (p *Playlist) Play() {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if p.CurrentSong == nil {
		return
	}

	if p.paused {
		p.paused = false
		p.currentTimer.Reset(p.CurrentSong.Value.(*Song).Duration)
		return
	}

	p.currentTimer = time.NewTimer(p.CurrentSong.Value.(*Song).Duration)
	go func() {
		<-p.currentTimer.C
		err := p.Next()
		if err != nil {
			fmt.Errorf("error with next song while playing %s", err.Error())
			return
		}
	}()
}

func (p *Playlist) Pause() {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if p.currentTimer != nil {
		p.currentTimer.Stop()
		p.paused = true
	}
}

func (p *Playlist) AddSong(newSong *Song) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if p.songs.Len() == 0 {
		p.songs.PushBack(newSong)
		p.CurrentSong = p.songs.Front()
		return
	}

	p.songs.PushBack(newSong)
}

func (p *Playlist) Next() error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if p.songs.Len() == 0 {
		return errors.New("playlist is empty")
	}

	if p.CurrentSong == nil {
		return errors.New("playlist is empty or current song is not available")
	}

	var next *list.Element

	if p.songs.Len() == 1 {
		next = p.songs.Front()
	} else {
		next = p.CurrentSong.Next()
	}

	if next != nil {
		p.CurrentSong = next
		p.Play()
	}
	return nil
}

func (p *Playlist) Prev() error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if p.songs.Len() == 0 {
		return errors.New("playlist is empty")
	}

	if p.CurrentSong == nil {
		return errors.New("playlist is empty or current song is not available")
	}

	if p.CurrentSong == nil {
		return nil
	}

	var prev *list.Element

	if p.songs.Len() == 1 {
		prev = p.songs.Front()
	} else {
		prev = p.CurrentSong.Prev()
	}

	if prev != nil {
		p.CurrentSong = prev
		p.Play()
	}
	return nil
}
