package middleware

import (
	"code94/internal/config"

	"github.com/gofiber/fiber/v2"
)

func ConfigMiddleware(cfg config.App) fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Locals("config", cfg)
		return c.Next()
	}
}
