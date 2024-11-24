package domain_test

import (
	"code94/internal/domain"
	"code94/internal/models"
	"code94/pkg/inmem"
	"context"
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func TestPostService_Create(t *testing.T) {
	ctx := context.Background()
	logger := zerolog.New(&testWriter{t: t})
	mockPostRepository := inmem.NewInMemoryRepository[models.Post]()
	container := domain.Container{
		PostRepository: mockPostRepository,
	}

	postService := domain.NewPostService(container, logger)
	post := &models.Post{Content: "Test content"}
	postService.Create(ctx, post)

	expectedPost := *post
	actualPost, _ := mockPostRepository.Get(post.Id)
	assert.Equal(t, expectedPost, actualPost, "Test valid call")
}

func TestPostService_AddLike(t *testing.T) {
	ctx := context.Background()
	logger := zerolog.New(&testWriter{t: t})
	mockPostRepository := inmem.NewInMemoryRepository[models.Post]()
	mockRactionRepository := inmem.NewInMemoryRepository[models.Reaction]()
	container := domain.Container{
		PostRepository:     mockPostRepository,
		ReactionRepository: mockRactionRepository,
	}

	postService := domain.NewPostService(container, logger)
	post := &models.Post{Content: "Test content"}
	postService.Create(ctx, post)
	postService.AddLike(ctx, post.Id)

	actualReaction, _ := mockRactionRepository.Get(uint(1))
	expectedPostId := post.Id
	assert.Equal(t, expectedPostId, actualReaction.PostId, "Test valid call")

	actualError := postService.AddLike(ctx, 2)
	expectedError := domain.ErrPostNotFound
	assert.Equal(t, expectedError, actualError, "Test invalid call")
}

func TestPostService_Comment(t *testing.T) {
	ctx := context.Background()
	logger := zerolog.New(&testWriter{t: t})
	mockPostRepository := inmem.NewInMemoryRepository[models.Post]()
	mockCommentRepository := inmem.NewInMemoryRepository[models.Comment]()
	container := domain.Container{
		PostRepository:    mockPostRepository,
		CommentRepository: mockCommentRepository,
	}

	postService := domain.NewPostService(container, logger)
	post := &models.Post{Content: "Test content"}
	postService.Create(ctx, post)
	postService.AddComment(ctx, post.Id, models.Comment{Content: "Test content"})

	actualComment, _ := mockCommentRepository.Get(uint(1))
	expectedPostId := post.Id
	assert.Equal(t, expectedPostId, actualComment.PostId, "Test valid call")

	actualError := postService.AddComment(ctx, 2, models.Comment{Content: "Test content"})
	expectedError := domain.ErrPostNotFound
	assert.Equal(t, expectedError, actualError, "Test invalid call")
}

func TestPostService_Update(t *testing.T) {
	ctx := context.Background()
	logger := zerolog.New(&testWriter{t: t})
	mockPostRepository := inmem.NewInMemoryRepository[models.Post]()
	container := domain.Container{
		PostRepository: mockPostRepository,
	}

	postService := domain.NewPostService(container, logger)
	post := &models.Post{Content: "Test content"}
	postService.Create(ctx, post)

	postToBeUpdated := models.Post{Content: "Test updated content"}
	postService.Update(ctx, 1, postToBeUpdated)

	expectedPost := postToBeUpdated
	actualPost, _ := mockPostRepository.Get(post.Id)
	assert.Equal(t, expectedPost.Content, actualPost.Content, "Test valid call")

	expectedError := domain.ErrPostNotFound
	actualError := postService.Update(ctx, 1, postToBeUpdated)

	assert.Equal(t, expectedError, actualError, "Test invalid call")
}

func TestPostService_RetrieveAll(t *testing.T) {
	ctx := context.Background()
	logger := zerolog.New(&testWriter{t: t})
	mockPostRepository := inmem.NewInMemoryRepository[models.Post]()
	container := domain.Container{
		PostRepository: mockPostRepository,
	}

	postService := domain.NewPostService(container, logger)
	post1 := &models.Post{Content: "Test content 1"}
	post2 := &models.Post{Content: "Test content 1"}

	postService.Create(ctx, post1)
	postService.Create(ctx, post2)

	actualCount := len(postService.RetriveAll(ctx))
	expectedCount := 2
	assert.Equal(t, expectedCount, actualCount, "Test valid call")

}

func TestPostService_Retrieve(t *testing.T) {
	ctx := context.Background()
	logger := zerolog.New(&testWriter{t: t})
	mockPostRepository := inmem.NewInMemoryRepository[models.Post]()
	container := domain.Container{
		PostRepository: mockPostRepository,
	}

	postService := domain.NewPostService(container, logger)
	post := &models.Post{Content: "Test content"}
	postService.Create(ctx, post)

	_, actualError := postService.Retrive(ctx, 2)
	expectedError := domain.ErrPostNotFound
	assert.Equal(t, expectedError, actualError, "Test invalid call")

}

type testWriter struct {
	t *testing.T
}

func (tw *testWriter) Write(p []byte) (n int, err error) {
	tw.t.Log(string(p))
	return len(p), nil
}
