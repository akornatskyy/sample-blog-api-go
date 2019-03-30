package search

import (
	"github.com/akornatskyy/sample-blog-api-go/posts/domain"
)

func Process(req *Request) (*Response, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	posts, err := domain.PostRepository().SearchPosts(
		req.Query, PageSize+1, req.Page*PageSize)
	if err != nil {
		return nil, err
	}
	paging := &Paging{}
	if req.Page > 0 {
		b := req.Page - 1
		paging.Before = &b
	}
	if len(posts) > PageSize {
		posts = posts[:PageSize]
		a := req.Page + 1
		paging.After = &a
	}
	resp := Response{
		Paging: paging,
		Items:  posts,
	}
	return &resp, nil
}
