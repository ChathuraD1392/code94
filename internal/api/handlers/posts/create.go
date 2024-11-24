package posts

import (
	"code94/internal/domain"
	"code94/internal/models"
	"code94/pkg/fiberutils"

	"github.com/gofiber/fiber/v2"
)

// CreateHandler handles the HTTP POST request to create a new post.
// It parses the request body to obtain the post data, validates it,
// and then attempts to create a new post using the post service.
// If the post is created successfully, it returns a success message with the post's ID;
// otherwise, it returns an error message indicating the failure.
func CreateHandler(ctr domain.Container) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		logger := fiberutils.GetLogger(ctx)
		// Create a post service instance using the container and logger.
		postService := domain.NewPostService(ctr, logger)

		// Parse the request body to get the post data.
		var post models.Post
		if err := ctx.BodyParser(&post); err != nil {
			logger.Error().Str("error", err.Error()).Msg("failed to parse request body.")
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "invalid request body",
			})
		}

		// Call the service to create the post.
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
