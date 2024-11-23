package api

import (
	"code94/internal/api/handlers/posts"
	"code94/internal/api/middleware"
	"code94/internal/config"
	"code94/internal/domain"
	"code94/pkg/fiberutils"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

func InitServer(cfg config.App, logger zerolog.Logger, cntr domain.Container) error {
	app := fiber.New()
	app.Use(
		fiberutils.LoggingMiddleware(logger),
		middleware.ConfigMiddleware(cfg))

	v1App := app.Group("/v1")

	v1App.Post("/posts", posts.CreateHandler(cntr))
	v1App.Put("/posts/:id", posts.UpdateHandler(cntr))
	v1App.Get("/posts", posts.RetrieveAllPostsHandler(cntr))
	v1App.Get("/posts/:id", posts.RetrievePostByIDHandler(cntr))

	v1App.Post("/posts/:id/like", posts.RactionHandler(cntr))
	v1App.Post("/posts/:id/comment", posts.CommnetHandler(cntr))

	addr := fmt.Sprintf(":%v", cfg.Server.Port)
	return app.Listen(addr)
}
