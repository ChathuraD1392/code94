package posts

import (
	"code94/internal/domain"
	"code94/pkg/fiberutils"

	"github.com/gofiber/fiber/v2"
)

func RetrievePostByIDHandler(ctr domain.Container) fiber.Handler {
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
