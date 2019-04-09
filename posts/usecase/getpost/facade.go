package getpost

import (
	"github.com/akornatskyy/sample-blog-api-go/posts/domain"
)

const maxCommentsAwaitingModeration = 5

func Process(req *Request) (*Response, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	p, err := domain.PostRepository().GetPost(req.Slug)
	if err != nil {
		return nil, ErrNotFound
	}
	resp := Response{
		Post: p,
	}

	return &resp, nil
}
