package addcomment

import (
	"github.com/akornatskyy/goext/errorstate"
	"github.com/akornatskyy/sample-blog-api-go/shared/rule"
)

func (req *Request) Validate() error {
	e := &errorstate.ErrorState{
		Domain: "posts",
	}

	rule.Slug.Validate(e, req.Slug)
	rule.Message.Validate(e, req.Message)

	return e.OrNil()
}
