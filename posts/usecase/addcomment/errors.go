package addcomment

import (
	"github.com/akornatskyy/goext/errorstate"
)

var (
	ErrNotFound = errorstate.Single(&errorstate.Detail{
		Domain:   "posts",
		Type:     "summary",
		Location: "post",
		Reason:   "not found",
		Message:  "We're sorry... the post cannot be found.",
	})

	ErrForbidden = errorstate.Single(&errorstate.Detail{
		Domain:   "posts",
		Type:     "summary",
		Location: "add post comment",
		Reason:   "forbidded",
		Message:  "Unauthorized to add comment to post.",
	})

	ErrTooManyComments = errorstate.Single(&errorstate.Detail{
		Domain:   "posts",
		Type:     "summary",
		Location: "add post comment",
		Reason:   "too many comments",
		Message:  "There are too many of your comments awaiting moderation. " +
			"Come back later, please.",
	})

	ErrFailed = errorstate.Single(&errorstate.Detail{
		Domain:   "posts",
		Type:     "summary",
		Location: "add post comment",
		Reason:   "failed",
		Message:  "We're sorry... the comment cannot be added.",
	})
)
