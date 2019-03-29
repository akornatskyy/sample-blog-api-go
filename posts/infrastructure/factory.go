package infrastructure

import (
	"github.com/akornatskyy/sample-blog-api-go/posts/domain"
	"github.com/akornatskyy/sample-blog-api-go/posts/domain/post"
	"github.com/akornatskyy/sample-blog-api-go/posts/infrastructure/mock"
)

type (
	factory struct {
	}
)

func NewFactory() domain.Factory {
	return factory{}
}

func (factory) PostRepository() post.Repository {
	return mock.NewPostRepository()
}
