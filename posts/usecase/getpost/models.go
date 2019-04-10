package getpost

import (
	"github.com/akornatskyy/sample-blog-api-go/posts/domain/post"
	"github.com/akornatskyy/sample-blog-api-go/shared/security"
)

type (
	Request struct {
		Slug      string
		Fields    string
		Principal security.Principal
	}

	Response struct {
		*post.Post

		Comments []*post.Comment `json:"comments,omitempty"`
	}
)
