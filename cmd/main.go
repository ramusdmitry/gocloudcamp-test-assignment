package main

import (
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	musicPlayer "music-player"
	"music-player/pkg/handler"
	"music-player/pkg/model"
	"music-player/pkg/repository"
	"music-player/pkg/service"
)

func main() {

	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	cfg, err := musicPlayer.LoadConfig("configs", "config")

	if err != nil {
		logrus.Fatalf("Failed to init config/envs: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(cfg.DB)

	playlist := model.NewPlaylist()

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services, playlist)

	server := new(musicPlayer.Server)
	if err := server.Run(cfg.Server.Port, handlers.InitRoutes()); err != nil {
		logrus.Fatalf("Error with running server: %s", err.Error())
	}
}
