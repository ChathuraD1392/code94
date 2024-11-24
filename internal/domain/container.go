package domain

import (
	"code94/internal/models"
	"code94/pkg/inmem"
)

// Container holds the repositories used across the application.
type Container struct {
	PostRepository     inmem.Repository[models.Post]
	ReactionRepository inmem.Repository[models.Reaction]
	CommentRepository  inmem.Repository[models.Comment]
}

// NewContainer initializes a new Container instance with the provided repositories.
func NewContainer(
	postRepository inmem.Repository[models.Post],
	reactionRepository inmem.Repository[models.Reaction],
	commentRepository inmem.Repository[models.Comment]) Container {
	return Container{
		PostRepository:     postRepository,
		ReactionRepository: reactionRepository,
		CommentRepository:  commentRepository,
	}
}
