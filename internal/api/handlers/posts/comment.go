package posts

import (
	"code94/internal/domain"
	"code94/internal/models"
	"code94/pkg/fiberutils"

	"github.com/gofiber/fiber/v2"
)

// CommnetHandler handles the HTTP POST request to add a comment to a post.
// It parses the request body to obtain the comment data, validates it,
// and then attempts to add the comment to the specified post using the post service.
// If the comment is added successfully, it returns a success message;
// otherwise, it returns an error message indicating the failure.
func CommnetHandler(ctr domain.Container) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		logger := fiberutils.GetLogger(ctx)
		// Create a post service instance using the container and logger.
		postService := domain.NewPostService(ctr, logger)

		// Get the post ID from the URL parameters.
		id, err := ctx.ParamsInt("id")
		if err != nil {
			logger.Error().Str("error", err.Error()).Msg("invalid ID parameter.")
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "invalid post ID",
			})
		}

		// Parse the request body to get the comment data.
		var commnet models.Comment
		if err := ctx.BodyParser(&commnet); err != nil {
			logger.Error().Str("error", err.Error()).Msg("failed to parse request body.")
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "invalid request body",
			})
		}

		// Call the service to add the comment to the post.
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

		// return sccess message
		return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message": "commented successfully",
		})
	}
}
