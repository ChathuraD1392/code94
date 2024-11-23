package fiberutils

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func GetLogger(ctx *fiber.Ctx) zerolog.Logger {
	defaultLogger := log.With().Logger()
	logger, ok := ctx.Locals("logger").(zerolog.Logger)
	if !ok {
		defaultLogger.Warn().Msg("Unbale to get the logger from the context")
		return defaultLogger
	}
	return logger
}

func LoggingMiddleware(logger zerolog.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		logger := log.With().Str("uuid", uuid.New().String()).Logger()
		c.Locals("logger", logger)
		return c.Next()
	}
}
