package posts

import (
	"code94/internal/domain"
	"code94/pkg/fiberutils"

	"github.com/gofiber/fiber/v2"
)

// RactionHandler handles the HTTP POST request to react to a post (like it).
// It parses the post ID from the URL parameters and adds a like reaction using the post service.
// If the post is found and the like is added successfully, it returns a success message;
// otherwise, it returns an error message based on the issue encountered.
func RactionHandler(ctr domain.Container) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		logger := fiberutils.GetLogger(ctx)
		// Create a post service instance using the container and logger.
		postService := domain.NewPostService(ctr, logger)

		// Parse the post ID from the URL parameters.
		id, err := ctx.ParamsInt("id")
		if err != nil {
			logger.Error().Str("error", err.Error()).Msg("invalid ID parameter.")
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "invalid post ID",
			})
		}

		// Call the service to add a like (reaction) to the post
		err = postService.AddLike(ctx.Context(), uint(id))
		if err != nil {
			if err == domain.ErrPostNotFound {
				return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
					"error": err.Error(),
				})
			}

			logger.Error().Str("error", err.Error()).Msg("unbale to react the post")
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "unbale to react the post",
			})
		}

		return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message": "Reacted successfully.",
		})

	}
}
