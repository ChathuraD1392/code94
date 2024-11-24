package domain

import (
	"code94/internal/models"
	"code94/pkg/inmem"
	"context"
	"sync"

	"github.com/rs/zerolog"
)

// PostService interface defines the methods for managing posts and their related actions.
type PostService interface {
	Create(ctx context.Context, post *models.Post) error
	Update(ctx context.Context, id uint, post models.Post) error
	RetriveAll(ctx context.Context) []models.Post
	Retrive(ctx context.Context, id uint) (models.DetailedPost, error)

	AddLike(ctx context.Context, postID uint) error
	AddComment(ctx context.Context, postID uint, comment models.Comment) error
}

// NewPostService initializes and returns an implementation of the PostService interface.
func NewPostService(cntr Container, logger zerolog.Logger) PostService {
	return &PostServiceImpl{
		cntr:   cntr,
		logger: logger,
	}
}

// PostServiceImpl is the concrete implementation of the PostService interface.
type PostServiceImpl struct {
	cntr   Container
	logger zerolog.Logger
}

// Create adds a new post to the PostRepository.
func (s *PostServiceImpl) Create(ctx context.Context, post *models.Post) error {
	return s.cntr.PostRepository.Add(post)
}

// Update modifies an existing post in the PostRepository. Returns an error if the post is not found.
func (s *PostServiceImpl) Update(ctx context.Context, id uint, post models.Post) error {
	err := s.cntr.PostRepository.Update(id, post)
	if err == inmem.ErrNotFound {
		return ErrPostNotFound
	}
	return err
}

// RetriveAll fetches all posts from the PostRepository.
func (s *PostServiceImpl) RetriveAll(ctx context.Context) []models.Post {
	return s.cntr.PostRepository.GetAll()
}

// Retrive fetches a specific post by ID, along with its comments and reactions.
func (s *PostServiceImpl) Retrive(ctx context.Context, id uint) (models.DetailedPost, error) {
	// Fetch the post by ID.
	post, err := s.cntr.PostRepository.Get(id)
	if err != nil {
		if err == inmem.ErrNotFound {
			return models.DetailedPost{}, ErrPostNotFound
		}
		return models.DetailedPost{}, err
	}

	// Use WaitGroup to fetch comments and reactions concurrently.
	var wg sync.WaitGroup
	var comments []models.Comment
	var reactions []models.Reaction
	wg.Add(2)

	// Fetch comments in a separate goroutine.
	go func() {
		defer wg.Done()
		comments = s.cntr.CommentRepository.Filter("PostId", post.Id)
	}()

	// Fetch reactions in a separate goroutine.
	go func() {
		defer wg.Done()
		reactions = s.cntr.ReactionRepository.Filter("PostId", post.Id)
	}()

	wg.Wait() // Wait for all goroutines to complete.

	// Combine post, comments, and reactions into a DetailedPost model.
	detailPost := models.DetailedPost{
		Id:        post.Id,
		Content:   post.Content,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
		Reactions: reactions,
		Comments:  comments,
	}

	return detailPost, nil
}

// AddLike adds a "like" reaction to a post. Returns an error if the post is not found.
func (s *PostServiceImpl) AddLike(ctx context.Context, postID uint) error {
	// Ensure the post exists.
	if _, err := s.cntr.PostRepository.Get(postID); err == inmem.ErrNotFound {
		return ErrPostNotFound
	}

	// Create a new "like" reaction and add it to the ReactionRepository.
	reaction := models.Reaction{
		Type:   models.Like,
		PostId: postID,
	}
	return s.cntr.ReactionRepository.Add(&reaction)
}

// AddComment adds a new comment to a post. Returns an error if the post is not found.
func (s *PostServiceImpl) AddComment(ctx context.Context, postID uint, comment models.Comment) error {
	// Ensure the post exists.
	if _, err := s.cntr.PostRepository.Get(postID); err == inmem.ErrNotFound {
		return ErrPostNotFound
	}

	// Assign the post ID to the comment and add it to the CommentRepository.
	comment.PostId = postID
	return s.cntr.CommentRepository.Add(&comment)
}
