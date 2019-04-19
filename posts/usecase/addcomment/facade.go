package addcomment

import (
	"github.com/akornatskyy/sample-blog-api-go/posts/domain"
	"github.com/akornatskyy/sample-blog-api-go/posts/usecase/getpost"
)

func Process(req *Request) error {
	if err := req.Validate(); err != nil {
		return err
	}
	if !req.Principal.IsAuthenticated() {
		return ErrForbidden
	}

	postID, err := domain.PostRepository().GetPostId(req.Slug)
	if err != nil {
		return ErrNotFound
	}

	n, err := domain.PostRepository().CountCommentsAwaitingModeration(
		req.Principal.ID, getpost.MaxCommentsAwaitingModeration)
	if err != nil || n >= getpost.MaxCommentsAwaitingModeration {
		return ErrTooManyComments
	}

	if err := domain.PostRepository().AddPostComment(
		postID, req.Principal.ID, req.Message); err != nil {
		return ErrFailed
	}

	return nil
}
