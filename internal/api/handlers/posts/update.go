package posts

import (
	"code94/internal/domain"
	"code94/internal/models"
	"code94/pkg/fiberutils"
	"code94/pkg/inmem"

	"github.com/gofiber/fiber/v2"
)

func UpdateHandler(ctr domain.Container) fiber.Handler {
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

		var post models.Post
		if err := ctx.BodyParser(&post); err != nil {
			logger.Error().Str("error", err.Error()).Msg("failed to parse request body.")
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "invalid request body",
			})
		}

		err = postService.Update(ctx.Context(), uint(id), post)
		if err != nil {
			if err == inmem.ErrNotFound {
				return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
					"error": "post not found",
				})
			}

			logger.Error().Str("error", err.Error()).Msg("failed to update post.")
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "failed to update post",
			})
		}

		// Return success response
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "post updated successfully",
		})
	}
}
