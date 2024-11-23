package posts

import (
	"code94/internal/domain"
	"code94/internal/models"
	"code94/pkg/fiberutils"

	"github.com/gofiber/fiber/v2"
)

func CreateHandler(ctr domain.Container) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		logger := fiberutils.GetLogger(ctx)
		postService := domain.NewPostService(ctr, logger)

		var post models.Post
		if err := ctx.BodyParser(&post); err != nil {
			logger.Error().Str("error", err.Error()).Msg("failed to parse request body.")
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "invalid request body",
			})
		}
		err := postService.Create(ctx.Context(), &post)
		if err != nil {
			logger.Error().Str("error", err.Error()).Msg("failed to create post")
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "failed to create post",
			})
		}

		// Return success response
		return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message": "post created successfully",
			"id":      post.Id,
		})
	}
}
