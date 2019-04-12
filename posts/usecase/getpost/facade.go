package getpost

import (
	"strings"
	"sync"

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

	var wg sync.WaitGroup
	if strings.Contains(req.Fields, "permissions") {
		resp.Permissions = &Permissions{}
		if req.Principal.IsAuthenticated() {
			wg.Add(1)
			go func() {
				defer wg.Done()
				n, err := domain.PostRepository().CountCommentsAwaitingModeration(
					req.Principal.ID, maxCommentsAwaitingModeration)
				if err == nil {
					resp.Permissions.CreateComment = n < maxCommentsAwaitingModeration
				}
			}()
		}
	}

	if strings.Contains(req.Fields, "comments") {
		wg.Add(1)
		go func() {
			defer wg.Done()
			comments, err := domain.PostRepository().ListComments(
				p.ID, req.Principal.ID)
			if err == nil {
				resp.Comments = comments
			}
		}()
	}

	wg.Wait()
	return &resp, nil
}
