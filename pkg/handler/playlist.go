package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"music-player/pkg/model"
	"net/http"
	"time"
)

type songInput struct {
	Title    string        `json:"title" binding:"required"`
	Duration time.Duration `json:"duration" binding:"required"`
}

func (h *Handler) getSongs(c *gin.Context) {
	go func() {
		songs, err := h.services.Player.GetSongs(1)

		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, "Не удалось получить треки", err.Error())
			return
		}
		c.JSON(http.StatusOK, map[string]interface{}{
			"status": "success",
			"data":   songs,
		})
	}()
}

func (h *Handler) pauseSong(c *gin.Context) {
	go func() {
		h.playlist.Pause()
	}()

	cur := h.playlist.CurrentSong

	var message string
	if cur == nil {
		message = "Ни одна композиция не была запущена"
	} else {
		message = fmt.Sprintf("Трек %s поставлен на паузу", cur.Value.(model.Song).Title) // *?
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": message,
	})
}

func (h *Handler) playSong(c *gin.Context) {
	go func() {
		h.playlist.Play()
	}()

	cur := h.playlist.CurrentSong

	var message string
	if cur == nil {
		message = "Ни одна композиция не была запущена"
	} else {
		message = fmt.Sprintf("Играет трек: %s", cur.Value.(model.Song).Title) // *?
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": message,
	})

}

func (h *Handler) prevSong(c *gin.Context) {

	go func() {
		err := h.playlist.Prev()
		if err != nil {
			newErrorResponse(c, http.StatusBadRequest, "Не удалось запустить предыдущий трек", err.Error())
			return
		}
	}()

	song := h.playlist.CurrentSong.Value.(*model.Song)

	c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": fmt.Sprintf("Предыдущая песня '%s' продолжительностью xx сек. %d", song.Title, song.Duration),
	})
}

func (h *Handler) nextSong(c *gin.Context) {

	go func() {
		err := h.playlist.Next()
		if err != nil {
			newErrorResponse(c, http.StatusBadRequest, "Не удалось переключить трек", err.Error())
			return
		}
	}()

	song := h.playlist.CurrentSong.Value.(*model.Song)

	//if nextSong == nil {
	//	newErrorResponse(c, http.StatusInternalServerError, "Трек недоступен", "")
	//	return
	//}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": fmt.Sprintf("Следующая песня '%s' продолжительностью xx сек. %d", song.Title, song.Duration),
	})
}

func (h *Handler) addSong(c *gin.Context) {

	var input model.Song

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Входные данные трека невалидные", err.Error())
		return
	}

	_, err := h.services.Player.AddSong(&input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Не удалось добавить песню в БД", err.Error())
		return
	}

	go func() {
		h.playlist.AddSong(&input)
	}()

	c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": fmt.Sprintf("Песня '%s' продолжительностью %d сек. добавлена", input.Title, input.Duration),
	})

}
