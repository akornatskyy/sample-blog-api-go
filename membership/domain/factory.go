package domain

import "github.com/akornatskyy/sample-blog-api-go/membership/domain/user"

type (
	Factory interface {
		UserRepository() user.Repository
	}
)

var singleton Factory

func GetFactory() Factory {
	return singleton
}

func SetFactory(f Factory) {
	singleton = f
}
