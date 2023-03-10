package handler

import (
	"github.com/gin-gonic/gin"
	"music-player/pkg/model"
	"music-player/pkg/service"
)

type Handler struct {
	services *service.Service
	playlist *model.Playlist
}

func NewHandler(services *service.Service, playlist *model.Playlist) *Handler {
	return &Handler{services: services, playlist: playlist}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		api.POST("/add", h.addSong)
		api.GET("/next", h.nextSong)
		api.GET("/prev", h.prevSong)
		api.GET("/play", h.playSong)
		api.GET("/pause", h.pauseSong)
		api.DELETE("/:id") // хендлеры для удаления и изменения
		api.PUT("/:id")    //
		api.GET("/", h.getSongs)
		api.GET("/get", h.getSongs)
	}

	return router
}
