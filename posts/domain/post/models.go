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

	Comment struct {
		Author    *Author   `json:"author"`
		Created   time.Time `json:"created_on"`
		Message   string    `json:"message"`
		Moderated bool      `json:"moderated"`
	}
)
