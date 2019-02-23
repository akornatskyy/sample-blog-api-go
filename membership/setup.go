package membership

import (
	"github.com/akornatskyy/sample-blog-api-go/membership/domain"
	"github.com/akornatskyy/sample-blog-api-go/membership/infrastructure"
	"github.com/akornatskyy/sample-blog-api-go/shared/config"
)

func Setup(c *config.Config) {
	domain.SetFactory(infrastructure.NewFactory())
	infrastructure.Routes(c)
}
