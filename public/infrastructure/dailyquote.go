package infrastructure

import (
	"net/http"

	"github.com/akornatskyy/goext/httpjson"
	"github.com/akornatskyy/sample-blog-api-go/public/domain"
)

func DailyQuoteHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := domain.QuoteRepository().FetchDailyQuote()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	httpjson.Encode(w, resp, http.StatusOK)
}
