package infrastructure

import (
	"github.com/akornatskyy/sample-blog-api-go/membership/domain"
	"github.com/akornatskyy/sample-blog-api-go/membership/domain/user"
	"github.com/akornatskyy/sample-blog-api-go/membership/infrastructure/mock"
)

type (
	factory struct {
	}
)

func Setup() {
	domain.SetFactory(NewFactory())
}

func NewFactory() domain.Factory {
	return factory{}
}

func (factory) UserRepository() user.Repository {
	return mock.NewUserRepository()
}
