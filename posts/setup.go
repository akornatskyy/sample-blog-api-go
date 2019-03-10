package posts

import (
	"github.com/akornatskyy/sample-blog-api-go/posts/infrastructure"
	"github.com/akornatskyy/sample-blog-api-go/shared/config"
)

func Setup(c *config.Config) {
	infrastructure.Routes(c)
}
