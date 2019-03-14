package search

type (
	Request struct {
		Query string `binding:"q"`
		Page  int    `binding:"page"`
	}

	Response struct {
	}
)
