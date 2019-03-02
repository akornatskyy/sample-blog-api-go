package quote

type Repository interface {
	FetchDailyQuote() (*Quote, error)
}
