package posts

import (
	"code94/internal/domain"
	"code94/pkg/fiberutils"

	"github.com/gofiber/fiber/v2"
)

// RetrievePostByIDHandler handles the HTTP GET request to retrieve a post by its ID.
// It parses the post ID from the URL parameters and retrieves the post details using the post service.
// If the post is found, it returns the post data; otherwise, an error message is returned.
func RetrievePostByIDHandler(ctr domain.Container) fiber.Handler {
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
		// Call the service to retrieve the post by ID.
		post, err := postService.Retrive(ctx.Context(), uint(id))
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
		return ctx.Status(fiber.StatusOK).JSON(post)
	}
}

// RetrieveAllPostsHandler handles the HTTP GET request to retrieve all posts.
// It calls the post service to get all posts and returns the posts in the response.
func RetrieveAllPostsHandler(ctr domain.Container) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		logger := fiberutils.GetLogger(ctx)
		postService := domain.NewPostService(ctr, logger)
		posts := postService.RetriveAll(ctx.Context())
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"posts": posts,
		})
	}
}
