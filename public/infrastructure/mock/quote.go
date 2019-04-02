package mock

import (
	"time"

	"github.com/akornatskyy/sample-blog-api-go/public/domain/quote"
	"github.com/akornatskyy/sample-blog-api-go/shared/mock"
)

type quoteRepository struct {
}

func NewQuoteRepository() quote.Repository {
	return &quoteRepository{}
}

func (*quoteRepository) FetchDailyQuote() (*quote.Quote, error) {
	return mock.DB.Quotes[time.Now().UTC().Day()%len(mock.DB.Quotes)], nil
}
