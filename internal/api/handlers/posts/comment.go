package posts

import (
	"code94/internal/domain"
	"code94/internal/models"
	"code94/pkg/fiberutils"

	"github.com/gofiber/fiber/v2"
)

func CommnetHandler(ctr domain.Container) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		logger := fiberutils.GetLogger(ctx)
		postService := domain.NewPostService(ctr, logger)

		id, err := ctx.ParamsInt("id")
		if err != nil {
			logger.Error().Str("error", err.Error()).Msg("invalid ID parameter.")
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "invalid post ID",
			})
		}

		var commnet models.Comment
		if err := ctx.BodyParser(&commnet); err != nil {
			logger.Error().Str("error", err.Error()).Msg("failed to parse request body.")
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "invalid request body",
			})
		}
		err = postService.AddComment(ctx.Context(), uint(id), commnet)
		if err != nil {
			if err == domain.ErrPostNotFound {
				return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
					"error": err.Error(),
				})
			}

			logger.Error().Str("error", err.Error()).Msg("failed to create comment")
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "failed to create comment",
			})
		}

		return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message": "commented successfully",
		})
	}
}
