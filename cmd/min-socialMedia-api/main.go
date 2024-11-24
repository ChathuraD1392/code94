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

// main function initializes the application by loading configurations,
// setting up repositories, and starting the server to handle API requests.
// It listens for OS signals (like interrupt) to gracefully shut down the application.

func main() {
	// Channel to receive OS signals (e.g., interrupt).
	osSig := make(chan os.Signal, 1)
	signal.Notify(osSig, os.Interrupt)

	// Create a new logger instance using zerolog.
	logger := zerolog.New(os.Stdout)

	// Define default configuration for the server.
	defaultConfig := config.App{
		Server: config.Server{
			Port: 8080, // Set the server port to 8080.
		},
	}

	// Load the configuration from the environment variable or default config.
	cfg, err := config.LoadConfig(configPathEnv, defaultConfig)
	if err != nil {
		// Log a warning if configuration loading fails.
		logger.Warn().Err(err).Msg("Unable to load configurations from MIN_SM_API_CONFIG_PATH")
	}

	// Initialize in-memory repositories for posts, reactions, and comments.
	postRepository := inmem.NewInMemoryRepository[models.Post]()
	reactionRepository := inmem.NewInMemoryRepository[models.Reaction]()
	commentRepository := inmem.NewInMemoryRepository[models.Comment]()

	// Create a new domain container to provide access to repositories and services.
	container := domain.NewContainer(postRepository, reactionRepository, commentRepository)

	// Start the API server in a separate goroutine.
	go func() {
		err := api.InitServer(cfg, logger, container)
		if err != nil {
			// Log a fatal error and stop the application if the server fails to start.
			logger.Fatal().Err(err)
		}
	}()

	// Wait for an OS interrupt signal to gracefully shut down the application.
	<-osSig
}
