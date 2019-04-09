package post

import "time"

type (
	Post struct {
		ID      string    `json:"-"`
		Slug    string    `json:"slug"`
		Title   string    `json:"title"`
		Created time.Time `json:"created_on"`
		Author  *Author   `json:"author"`
		Message string    `json:"message"`
	}

	Author struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	}
)
