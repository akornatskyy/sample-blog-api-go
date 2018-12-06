package membership

import (
	"github.com/akornatskyy/sample-blog-api-go/membership/domain"
	"github.com/akornatskyy/sample-blog-api-go/membership/infrastructure"
)

func Setup() {
	domain.SetFactory(infrastructure.NewFactory())
	infrastructure.Routes()
}
