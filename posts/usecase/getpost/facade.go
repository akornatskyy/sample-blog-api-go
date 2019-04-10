package getpost

import (
	"strings"

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

	if strings.Contains(req.Fields, "comments") {
		comments, err := domain.PostRepository().ListComments(
			p.ID, req.Principal.ID)
		if err == nil {
			resp.Comments = comments
		}
	}

	return &resp, nil
}
