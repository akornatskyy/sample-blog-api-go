package mock

import (
	"log"
	"time"

	"github.com/akornatskyy/sample-blog-api-go/public/domain/quote"
	"github.com/akornatskyy/sample-blog-api-go/shared/iojson"
)

type quoteRepository struct {
}

var (
	quotes = loadQuotes("quote-samples.json")
)

func NewQuoteRepository() quote.Repository {
	return quoteRepository{}
}

func (r quoteRepository) FetchDailyQuote() (*quote.Quote, error) {
	return &quotes[time.Now().UTC().Day()%len(quotes)], nil
}

func loadQuotes(filename string) []quote.Quote {
	var r struct {
		Quotes []quote.Quote `json:"quotes"`
	}
	if err := iojson.ReadFile(filename, &r); err != nil {
		panic(err)
	}

	log.Printf("loaded %d quotes", len(r.Quotes))

	return r.Quotes
}
