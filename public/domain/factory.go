package domain

import "github.com/akornatskyy/sample-blog-api-go/public/domain/quote"

type (
	Factory interface {
		QuoteRepository() quote.Repository
	}
)

var singleton Factory

func SetFactory(f Factory) {
	singleton = f
}

func QuoteRepository() quote.Repository {
	return singleton.QuoteRepository()
}
