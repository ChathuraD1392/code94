package domain

import (
	"code94/internal/models"
	"code94/pkg/inmem"
	"context"
	"sync"

	"github.com/rs/zerolog"
)

type PostService interface {
	Create(ctx context.Context, post *models.Post) error
	Update(ctx context.Context, id uint, post models.Post) error
	RetriveAll(ctx context.Context) []models.Post
	Retrive(ctx context.Context, id uint) (models.DetailedPost, error)

	AddLike(ctx context.Context, postID uint) error
	AddComment(ctx context.Context, postID uint, comment models.Comment) error
}

func NewPostService(cntr Container, logger zerolog.Logger) PostService {
	return &PostServiceImpl{
		cntr:   cntr,
		logger: logger,
	}
}

type PostServiceImpl struct {
	cntr   Container
	logger zerolog.Logger
}

func (s *PostServiceImpl) Create(ctx context.Context, post *models.Post) error {
	return s.cntr.PostRepository.Add(post)
}

func (s *PostServiceImpl) Update(ctx context.Context, id uint, post models.Post) error {
	err := s.cntr.PostRepository.Update(id, post)
	if err == inmem.ErrNotFound {
		return ErrPostNotFound
	}
	return err
}

func (s *PostServiceImpl) RetriveAll(ctx context.Context) []models.Post {
	return s.cntr.PostRepository.GetAll()
}

func (s *PostServiceImpl) Retrive(ctx context.Context, id uint) (models.DetailedPost, error) {
	post, err := s.cntr.PostRepository.Get(id)
	if err != nil {
		if err == inmem.ErrNotFound {
			return models.DetailedPost{}, ErrPostNotFound
		}
		return models.DetailedPost{}, err
	}
	var wg sync.WaitGroup
	var comments []models.Comment
	var reactions []models.Reaction
	wg.Add(2)

	go func() {
		defer wg.Done()
		comments = s.cntr.CommentRepository.Filter("PostId", post.Id)
	}()

	go func() {
		defer wg.Done()
		reactions = s.cntr.ReactionRepository.Filter("PostId", post.Id)
	}()

	wg.Wait()

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

func (s *PostServiceImpl) AddLike(ctx context.Context, postID uint) error {
	if _, err := s.cntr.PostRepository.Get(postID); err == inmem.ErrNotFound {
		return ErrPostNotFound
	}

	reaction := models.Reaction{
		Type:   models.Like,
		PostId: postID,
	}
	return s.cntr.ReactionRepository.Add(&reaction)
}

func (s *PostServiceImpl) AddComment(ctx context.Context, postID uint, comment models.Comment) error {
	if _, err := s.cntr.PostRepository.Get(postID); err == inmem.ErrNotFound {
		return ErrPostNotFound
	}

	comment.PostId = postID
	return s.cntr.CommentRepository.Add(&comment)
}
