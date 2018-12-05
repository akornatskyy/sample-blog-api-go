package membership

import (
	"github.com/akornatskyy/sample-blog-api-go/membership/domain"
	"github.com/akornatskyy/sample-blog-api-go/membership/infrastructure"
	"github.com/akornatskyy/sample-blog-api-go/shared/httptoken"
)

func Setup(t httptoken.Token) {
	domain.SetFactory(infrastructure.NewFactory())
}
