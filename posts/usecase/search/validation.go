package search

import (
	"github.com/akornatskyy/sample-blog-api-go/shared/errorstate"
	"github.com/akornatskyy/sample-blog-api-go/shared/rule"
)

func (req *Request) Validate() error {
	e := &errorstate.ErrorState{
		Domain: "posts",
	}

	rule.Query.Validate(e, req.Query)
	rule.Page.Validate(e, req.Page)

	return e.OrNil()
}
