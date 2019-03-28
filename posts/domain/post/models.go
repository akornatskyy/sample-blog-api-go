package post

import "time"

type (
	Post struct {
		Slug    string    `json:"slug"`
		Title   string    `json:"title"`
		Message string    `json:"message"`
		Created time.Time `json:"created_on"`
		Author  Author    `json:"author"`
	}

	Author struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	}
)
