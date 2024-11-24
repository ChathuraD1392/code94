package main

import (
	"code94/internal/api"
	"code94/internal/config"
	"code94/internal/domain"
	"code94/internal/models"
	"code94/pkg/inmem"
	"os"
	"os/signal"

	"github.com/rs/zerolog"
)

const configPathEnv = "MIN_SM_API_CONFIG_PATH"

func main() {

	osSig := make(chan os.Signal, 1)
	signal.Notify(osSig, os.Interrupt)

	logger := zerolog.New(os.Stdout)
	defaultConfig := config.App{
		Server: config.Server{
			Port: 8080,
		},
	}
	cfg, err := config.LoadConfig(configPathEnv, defaultConfig)
	if err != nil {
		logger.Warn().Err(err).Msg("Unable to load configurations from MIN_SM_API_CONFIG_PATH")
	}

	postRepository := inmem.NewInMemoryRepository[models.Post]()
	reactionRepository := inmem.NewInMemoryRepository[models.Reaction]()
	commentRepository := inmem.NewInMemoryRepository[models.Comment]()
	container := domain.NewContainer(postRepository, reactionRepository, commentRepository)
	go func() {
		err := api.InitServer(cfg, logger, container)
		if err != nil {
			logger.Fatal().Err(err)
		}
	}()

	<-osSig
}
