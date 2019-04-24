package infrastructure

import (
	"log"

	"github.com/akornatskyy/sample-blog-api-go/membership/domain"
	"github.com/akornatskyy/sample-blog-api-go/membership/domain/user"
	"github.com/akornatskyy/sample-blog-api-go/membership/infrastructure/mock"
	"github.com/akornatskyy/sample-blog-api-go/membership/infrastructure/sql"
	"github.com/akornatskyy/sample-blog-api-go/shared/config"
)

type (
	factory struct {
		user user.Repository
	}
)

func NewFactory(c *config.Config) domain.Factory {
	switch c.Strategy {
	case config.StrategySQL:
		return &factory{
			user: sql.NewUserRepository(c.DB),
		}
	case config.StrategyMock:
		return &factory{
			user: mock.NewUserRepository(),
		}
	}
	log.Fatal("unknown repository strategy")
	return nil
}

func (f *factory) UserRepository() user.Repository {
	return f.user
}
