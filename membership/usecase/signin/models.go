package signin

type (
	Request struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	Response struct {
		UserID int
	}
)
