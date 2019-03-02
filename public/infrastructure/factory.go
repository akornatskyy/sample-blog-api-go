package infrastructure

import (
	"github.com/akornatskyy/sample-blog-api-go/public/domain"
	"github.com/akornatskyy/sample-blog-api-go/public/domain/quote"
	"github.com/akornatskyy/sample-blog-api-go/public/infrastructure/mock"
)

type (
	factory struct {
	}
)

func NewFactory() domain.Factory {
	return factory{}
}

func (factory) QuoteRepository() quote.Repository {
	return mock.NewQuoteRepository()
}
