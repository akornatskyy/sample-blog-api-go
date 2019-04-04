package infrastructure

import (
	"log"

	"github.com/akornatskyy/sample-blog-api-go/posts/domain"
	"github.com/akornatskyy/sample-blog-api-go/posts/domain/post"
	"github.com/akornatskyy/sample-blog-api-go/posts/infrastructure/mock"
	"github.com/akornatskyy/sample-blog-api-go/shared/config"
)

type (
	factory struct {
		post post.Repository
	}
)

func NewFactory(c *config.Config) domain.Factory {
	switch c.Strategy {
	case config.StrategyMock:
		return &factory{
			post: mock.NewPostRepository(),
		}
	}
	log.Fatal("unknown repository strategy")
	return nil
}

func (f *factory) PostRepository() post.Repository {
	return f.post
}
