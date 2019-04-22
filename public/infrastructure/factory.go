package infrastructure

import (
	"log"

	"github.com/akornatskyy/sample-blog-api-go/public/domain"
	"github.com/akornatskyy/sample-blog-api-go/public/domain/quote"
	"github.com/akornatskyy/sample-blog-api-go/public/infrastructure/mock"
	"github.com/akornatskyy/sample-blog-api-go/shared/config"
)

type (
	factory struct {
		quote quote.Repository
	}
)

func NewFactory(c *config.Config) domain.Factory {
	switch c.Strategy {
	case config.StrategySQL:
		fallthrough
	case config.StrategyMock:
		return &factory{
			quote: mock.NewQuoteRepository(),
		}
	}
	log.Fatal("unknown repository strategy")
	return nil
}

func (f *factory) QuoteRepository() quote.Repository {
	return f.quote
}
