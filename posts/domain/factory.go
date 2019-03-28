package domain

import "github.com/akornatskyy/sample-blog-api-go/posts/domain/post"

type (
	Factory interface {
		PostRepository() post.Repository
	}
)

var singleton Factory

func SetFactory(f Factory) {
	singleton = f
}

func PostRepository() post.Repository {
	return singleton.PostRepository()
}
