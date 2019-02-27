package domain

import "github.com/akornatskyy/sample-blog-api-go/membership/domain/user"

type (
	Factory interface {
		UserRepository() user.Repository
	}
)

var singleton Factory

func SetFactory(f Factory) {
	singleton = f
}

func UserRepository() user.Repository {
	return singleton.UserRepository()
}
