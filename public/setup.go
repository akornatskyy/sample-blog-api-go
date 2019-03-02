package public

import (
	"github.com/akornatskyy/sample-blog-api-go/public/domain"
	"github.com/akornatskyy/sample-blog-api-go/public/infrastructure"
	"github.com/akornatskyy/sample-blog-api-go/shared/config"
)

func Setup(c *config.Config) {
	domain.SetFactory(infrastructure.NewFactory())
	infrastructure.Routes(c)
}
