package infrastructure

import (
	"github.com/akornatskyy/sample-blog-api-go/shared/config"
)

func Routes(c *config.Config) {
	c.Router.HandlerFunc("GET", "/quote/daily", DailyQuoteHandler)
}
