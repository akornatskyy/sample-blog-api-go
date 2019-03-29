package posts

import (
	"github.com/akornatskyy/sample-blog-api-go/posts/domain"
	"github.com/akornatskyy/sample-blog-api-go/posts/infrastructure"
	"github.com/akornatskyy/sample-blog-api-go/shared/config"
)

func Setup(c *config.Config) {
	domain.SetFactory(infrastructure.NewFactory())
	infrastructure.Routes(c)
}
